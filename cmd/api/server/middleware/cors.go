package middleware

import (
	"net/http"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
)

func CORSMiddleware(serverAuthClientsURLs []string) RequestMiddleware {
	return cors.NewCorsHandler(cors.Options{
		AllowedOrigins:   serverAuthClientsURLs,
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
		AllowCredentials: false,
		AllowMaxAge:      5600,
	}).CorsMiddleware
}
