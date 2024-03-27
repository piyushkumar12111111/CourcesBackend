package middleware

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        splitToken := strings.Split(authHeader, "Bearer ")
        if len(splitToken) != 2 {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        tokenString := splitToken[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return []byte("your_secret_key"), nil
        })

        if err != nil || !token.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
