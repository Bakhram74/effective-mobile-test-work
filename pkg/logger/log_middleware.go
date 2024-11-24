package logger

import (
	"log/slog"
	"net/http"
	"time"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (lw *logResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := newLoggingResponseWriter(w)

		next.ServeHTTP(lw, r)

		slog.Info("incoming request",
			slog.String("url", r.URL.RequestURI()),
			slog.String("user_agent", r.UserAgent()),
			slog.Duration("elapsed_ms", time.Since(start)),
			slog.Int("status_code", lw.statusCode),
			slog.String("method", r.Method))
	})
}
