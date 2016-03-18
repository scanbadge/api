package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/users"
	"log"
	"strings"
	"time"
)

var loginKey []byte

// GetUserID gets the user ID of the current request. It will return 0 and an error when something went wrong.
func GetUserID(c *gin.Context) (int64, error) {
	var encodedToken = c.Request.Header.Get("Authorization")

	token, err := ParseToken(encodedToken)

	if err == nil {
		if token.Valid {
			id := token.Claims["id"].(float64) // Required to prevent runtime panic due to wrong type assertion

			return int64(id), nil
		}
	}

	return 0, fmt.Errorf("Cannot get user ID from JWT due to %s", err.Error())
}

// ParseToken parses the encoded token by verifying if it contains 'Bearer {token}'. The token will be validated as well.
func ParseToken(encodedToken string) (*jwt.Token, error) {
	// Requirements of JWT token: not empty, longer than 6 characters, and must contain 'BEARER '.
	if encodedToken != "" && len(encodedToken) > 6 && strings.ToUpper(encodedToken[0:7]) == "BEARER " {
		encodedToken = encodedToken[7:]

		token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what we expect:
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
			log.Println("token is not in required format", err)
		}

		return token, nil
	}

	return nil, fmt.Errorf("invalid token provided")
}

func generateToken(user users.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set headers
	token.Header["kid"] = "login"
	// Set claims
	token.Claims["id"] = user.ID
	token.Claims["name"] = fmt.Sprintf("%s %s", user.Firstname, user.Lastname)
	token.Claims["admin"] = user.IsAdmin
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Sign and get the complete encoded token as a string
	key, err := lookupKey("login")

	if err == nil && key != nil {
		tokenString, err := token.SignedString(key)

		if err != nil {
			return "", fmt.Errorf("cannot create signed string for new JWT: %s", err)
		}

		return tokenString, nil
	}

	return "", err
}

func lookupKey(kind interface{}) (interface{}, error) {
	if str, ok := kind.(string); ok {
		switch str {
		case "login":
			return configuration.JwtKey, nil
		}
	}

	return nil, fmt.Errorf("unknown jwt kind")
}
