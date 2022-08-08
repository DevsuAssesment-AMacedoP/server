package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func HandleDevOps(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		devOpsPost(w, r)

	default:
		http.Error(w, "ERROR", http.StatusNotFound)
	}
}

func devOpsPost(w http.ResponseWriter, r *http.Request) {
	var req requestJson

	err := decodeJsonBody(w, r, &req)
	if err != nil {
		var er *errorResponse
		if errors.As(err, &er) {
			returnErrorResponse(w, *er)
		} else {
			returnErrorResponse(w, errorResponse{
				status: http.StatusInternalServerError,
				msg:    http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	resp := responseJson{Message: fmt.Sprintf("Hello %s your message will be sent", req.To)}

	returnOkResponse(w, resp)
}

func HandleJWT(jwtSigningKey []byte) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Don't validate, just return an accessToken
			expirationTime := time.Now().Add(15 * time.Minute)

			claims := jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(jwtSigningKey)
			if err != nil {
				fmt.Printf("%e", err)
				returnErrorResponse(w, errorResponse{
					status: http.StatusInternalServerError,
					msg:    "",
				})
				return
			}

			resp := accessTokenJSson{
				AccessToken: tokenString,
			}

			returnOkResponse(w, resp)

		default:
			returnErrorResponse(w, errorResponse{
				status: http.StatusNotFound,
				msg:    "",
			})
		}
	}
}
