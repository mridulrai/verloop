package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/mridulrai/verloop/pkg/logger"
)

// recoverHandler recovers the application in case of a panic.
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.InitSugarLogger().Panicw("Panic ",
					// Structured context as loosely typed key-value pairs.
					"error", err,
				)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// loggingHandler logs basic information about request.
func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger.InitSugarLogger().Infow("Request details",
			// Structured context as loosely typed key-value pairs.
			"requestID", r.Header.Get("X-Request-ID"),
			"method", r.Method,
			"url", r.URL.String(),
		)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// identifierHandler adds a unique id in header for each request if not present.
func identifierHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("X-Request-ID")) == 0 {
			r.Header.Set("X-Request-ID", uuid.New().String())
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
