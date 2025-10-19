package database

import (
	"context"
	"net/http"
	"time"
)

const ContextKey = "database"

func TransactionMiddleware(db *Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			gorm := db.Db()
			tx := gorm.WithContext(ctx).Begin()
			defer func() {
				if rec := recover(); rec != nil {
					tx.Rollback()
					panic(rec)
				}
			}()

			r = r.WithContext(context.WithValue(ctx, ContextKey, tx))
			rr := &statusRecorder{ResponseWriter: w}
			next.ServeHTTP(rr, r)

			if rr.status >= 400 {
				tx.Rollback()
				return
			}

			if err := tx.Commit().Error; err != nil {
				http.Error(w, "transaction commit failed", http.StatusInternalServerError)
			}
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}
