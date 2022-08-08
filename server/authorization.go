package server

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ApiKeyMiddleware(apiKey string) func(handler http.Handler) http.Handler {
	apiKeyHeader := "X-Parse-REST-API-Key"

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKeyRequest := r.Header.Get(apiKeyHeader)
			if apiKeyRequest == "" {
				returnErrorResponse(w, errorResponse{
					status: http.StatusUnauthorized,
					msg:    "No API Key set",
				})
				return
			}

			if apiKeyRequest != apiKey {
				returnErrorResponse(w, errorResponse{
					status: http.StatusUnauthorized,
					msg:    "Invalid API Key",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func JWTMiddleware(jwtSinginKey []byte) func(handler http.Handler) http.Handler {
	jwtHeader := "X-JWT-KWY"

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtRequest := r.Header.Get(jwtHeader)

			if jwtRequest == "" {
				returnErrorResponse(w, errorResponse{
					status: http.StatusUnauthorized,
					msg:    "No JWT set",
				})
				return
			}

			// Just validate signature and expiration
			_, err := jwt.Parse(jwtRequest, func(t *jwt.Token) (interface{}, error) {
				return jwtSinginKey, nil
			}, jwt.WithValidMethods(jwt.GetAlgorithms()))

			if err != nil {
				returnErrorResponse(w, errorResponse{
					status: http.StatusUnauthorized,
					msg:    fmt.Sprintf("%v", err),
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
