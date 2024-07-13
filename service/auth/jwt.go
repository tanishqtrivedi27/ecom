package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tanishqtrivedi27/ecom/config"
	"github.com/tanishqtrivedi27/ecom/utils"
)

type contextKey string

const UserKey contextKey = "userId"

func CreateJWT(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiresIn": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// Auth middleware
func JWTAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token %v", err)
			permissionDenied(w)
			return
		}


		if !token.Valid {
			log.Printf("invalid token")
			permissionDenied(w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if (!ok) {
			log.Printf("invalid token claims")
			permissionDenied(w)
			return
		}

		if expiresIn, ok := claims["expiresIn"].(float64); ok {
			expirationTime := time.Unix(int64(expiresIn), 0)
			if time.Now().After(expirationTime) {
				log.Printf("token has expired")
				permissionDenied(w)
				return
			}
		} else {
			log.Printf("expiration time not found")
			permissionDenied(w)
			return
		}


		userIdStr, ok := claims["userId"].(string)
		if !ok {
			log.Printf("userId not found in token claims")
			permissionDenied(w)
			return
		}

		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			log.Printf("invalid userId in token claims")
			permissionDenied(w)
			return
		}

		// u, err := store.GetUserById(userId)
		// if err != nil {
		// 	log.Printf("failed to get user by id: %v", err)
		// 	permissionDenied(w)
		// 	return
		// }

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	token := strings.TrimPrefix(tokenAuth, "Bearer ")
	return token
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserFromContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userId
}
