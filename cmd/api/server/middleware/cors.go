package middleware

import (
	"net/http"

	"github.com/Nhanderu/fasthttpcors"
)

func CORSMiddleware(serverAuthClientsURLs []string) RequestMiddleware {
	return fasthttpcors.NewCorsHandler(fasthttpcors.Options{
		AllowedOrigins:   serverAuthClientsURLs,
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: false,
		AllowMaxAge:      60 * 60 * 24 * 30,
	}).CorsMiddleware
}
