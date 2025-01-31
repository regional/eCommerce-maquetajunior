package handlers

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ResolveClaims(w http.ResponseWriter, r *http.Request, claimKey string) (interface{}, error) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return "nil", http.ErrAbortHandler
	}
	tokenString := strings.ReplaceAll(bearerToken, "Bearer ", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma y devolver la clave de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte("my_secret_key"), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return "nil", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return "nil", http.ErrAbortHandler
	}
	
	value, ok := claims[claimKey]
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return "nil", http.ErrAbortHandler
	}

	return value, nil
}
