package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/users"
	"github.com/scanbadge/api/utility"
	"log"
	"time"
)

func isValidToken(encodedToken string) (success bool) {
	// If we do not check for an empty token, the token parsing will fail.
	if encodedToken == "" {
		log.Println(fmt.Errorf("Provided token is empty"))
		return false
	}

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		k, err := lookupKey(token.Header["kid"])
		if err != nil {
			return nil, err
		}

		return k, nil
	})

	if err != nil {
		checkErr("Could not parse token:", err)
	}

	return token.Valid
}

func generateToken(user users.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set headers
	token.Header["kid"] = "login"
	// Set claims
	token.Claims["sub"] = user.ID
	token.Claims["name"] = fmt.Sprintf("%s:%s", user.Firstname, user.Lastname)
	token.Claims["admin"] = user.IsAdmin
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	key, err := lookupKey("login")

	if err == nil && key != nil {
		tokenString, err := token.SignedString(key)

		if err != nil {
			return "", fmt.Errorf("Cannot create signed string for new JWT: %s", err)
		}

		return tokenString, nil
	}

	return "", err
}

func lookupKey(kind interface{}) (interface{}, error) {
	if str, ok := kind.(string); ok {
		switch str {
		case "login":
			cfile := configuration.Config.Key

			if cfile != "" {
				f, err := utility.ReadData(configuration.Config.Key)

				if err != nil {
					return nil, err
				}

				return utility.DecodeBase64(f)
			}

			break
		}
	}

	return nil, fmt.Errorf("unknown jwt kind")
}
