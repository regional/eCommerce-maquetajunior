package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func JWTAuthMiddleware(next http.Handler, allowdRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extraer el rol de los claims
		role, ok := (*claims)["rolename"].(string)
		if !ok {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		roleAllowed := false
		for _, allowedRole := range allowdRoles {
            if role == allowedRole {
                roleAllowed = true
                break
            }
		}

        if !roleAllowed {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

		next.ServeHTTP(w, r)
	})
}
