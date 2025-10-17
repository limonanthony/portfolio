package logger

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter

	status      int
	written     int
	wroteHeader bool
}

const requestIDHeader = "X-Request-ID"

func (rw *responseRecorder) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.status = code
		rw.wroteHeader = true
	}

	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseRecorder) Write(b []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}

	n, err := rw.ResponseWriter.Write(b)
	rw.written += n

	return n, err
}

func newReqID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])

	return hex.EncodeToString(b[:])
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get(requestIDHeader)
		if reqID == "" {
			reqID = newReqID()
		}

		w.Header().Set(requestIDHeader, reqID)

		start := time.Now()
		method := r.Method
		path := r.URL.Path
		proto := r.Proto
		cl := r.ContentLength

		Infof("request_start request_id=%s method=%s path=%s proto=%s content_length=%d",
			reqID, method, path, proto, cl)

		rr := &responseRecorder{ResponseWriter: w, status: 0}

		next.ServeHTTP(rr, r)

		status := rr.status
		if status == 0 {
			status = http.StatusOK
		}

		dur := time.Since(start)
		ms := float64(dur.Microseconds()) / 1000.0

		switch {
		case status >= http.StatusInternalServerError:
			Errorf("request_end request_id=%s method=%s path=%s status=%d bytes=%d duration_ms=%.3f",
				reqID, method, path, status, rr.written, ms)
		case status >= http.StatusBadRequest:
			Warnf("request_end request_id=%s method=%s path=%s status=%d bytes=%d duration_ms=%.3f",
				reqID, method, path, status, rr.written, ms)
		default:
			Successf("request_end request_id=%s method=%s path=%s status=%d bytes=%d duration_ms=%.3f",
				reqID, method, path, status, rr.written, ms)
		}
	})
}
