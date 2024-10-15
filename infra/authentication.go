package infra

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	SigningKey = []byte("allYourBase")
)

func SecureRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authoriztionHeader := r.Header.Get("Authorization")
		if authoriztionHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Missing a token header")
			return
		}
		token := strings.Split(authoriztionHeader, " ")[1]

		jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return SigningKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Missing a token: "+err.Error())
			return
		}

		if !jwtToken.Valid {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "the token is not valid ")
			return
		}

		next(w, r)
	}
}
