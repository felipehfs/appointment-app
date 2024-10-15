package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/felipehfs/appointment-app/infra"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct{}

func (a AuthController) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /token", a.GenerateToken)
}

func NewAuthController() AuthController {
	return AuthController{}
}

func (a AuthController) GenerateToken(w http.ResponseWriter, r *http.Request) {
	claims := &jwt.RegisteredClaims{
		Issuer:    "test",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(infra.SigningKey)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred trying to generate token: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": ss,
	})
}
