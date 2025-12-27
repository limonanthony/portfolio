package database

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

const ContextKey = "database"

func Middleware(db *Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), ContextKey, db.Db()))
			next.ServeHTTP(w, r)
		})
	}
}

func TransactionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbVal := r.Context().Value(ContextKey)
		if dbVal == nil {
			next.ServeHTTP(w, r)
			return
		}

		gormDB, ok := dbVal.(*gorm.DB)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		tx := gormDB.WithContext(ctx).Begin()

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

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}
