package middleware

import (
    "net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the session is valid
        session, err := store.Get(r, "session-name")
        if err != nil || session.Values["authenticated"] != true {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        // If valid, proceed to the next handler
        next.ServeHTTP(w, r)
    })
}