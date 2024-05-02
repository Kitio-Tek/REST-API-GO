package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc,store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the request (Auth header)
		tokenString := GetTokenFromRequest(r)
		// validate the token
		token,err := validateJWT(tokenString)
		if err != nil {
			// handle error
			log.Println ("failed to authenticate token,", err)

			permissionDenied(w)
			return
		}

		if !token.Valid {
			// handle error
			log.Println ("failed to authenticate token")

			permissionDenied(w)
			return
		}
		// get the userID from the token

		claims:=token.Claims.(jwt.MapClaims)
		userID:=claims["user_id"].(string)

		_,err=store.GetUserByID(userID)
		if err != nil {
			// handle error
			log.Println ("failed to get user", err)

			permissionDenied(w)
			return

		}
		// call the handler func and continue to the endpoint
		handlerFunc (w,r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, ErrorResponse{Error: "permission denied"})
}
func GetTokenFromRequest(r *http.Request) string {
	// code

	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != ""  {
		return tokenAuth  
	}
	if tokenQuery != "" {
		return tokenQuery
	}
	return ""
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	// code
	secret:=Envs.JWTSecret
	
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		
		return []byte(secret), nil
	})
}

func HashPassword(password string) (string,error) {
	// code
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		log.Println("failed to hash password", err)
		return "", err
	}
	return string(hash), nil
}

func CreateJWT(secret []byte, userID int64) (string, error) {
	// code
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("failed to create JWT", err)
		return "", err
	}
	
	return tokenString, nil
}