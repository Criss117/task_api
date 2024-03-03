package auth_middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	mysql "unicauca.edu.co/cristian/task-api/src/db"
	user_entity "unicauca.edu.co/cristian/task-api/src/user/entities"
)
type contextKey string
const LoggedInUser contextKey = "loggedInUser"

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
    	return
    }

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])	
			}

			return []byte(os.Getenv("SECRETKEY")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}
		
		exp := claims["exp"].(float64)
		expirationTime := time.Unix(int64(exp), 0)

		if time.Now().After(expirationTime) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}

		var user user_entity.User
		mysql.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid credentials"))
			return
		}

		ctx := context.WithValue(r.Context(), LoggedInUser, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}