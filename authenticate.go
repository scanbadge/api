package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/users"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// AuthRequired is middleware that validates whether or not the current request has been authenticated.
// If the request has not been authenticated, a 401 HTTP code will be returned.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/v1/auth" {
			// Allow user to obtain authentication token.
			c.Next()
		} else {
			// Get authorization header and trim ' Bearer '.
			var token = c.Request.Header.Get("Authorization")
			token = strings.TrimSpace(token)
			token = strings.TrimPrefix(token, "Bearer")
			token = strings.TrimSpace(token)

			if isValidToken(token) {
				// Authentication is OK, proceed.
				c.Next()
			} else {
				// Authentication is wrong, stop.
				c.JSON(401, gin.H{"error": "Unauthorized: Access is denied"})
				c.Abort()
			}
		}
	}
}

// Authenticate authenticates a user by his eail and password.
func Authenticate(c *gin.Context) {
	var user users.User
	err := c.BindJSON(&user)

	if err == nil && user.Username != "" && user.Password != "" {
		var selectedUser users.User
		err := configuration.Dbmap.SelectOne(&selectedUser, "select * from users where username = ?", user.Username)

		if err == nil && selectedUser.Password != "" {
			err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
			if err == nil {
				//
				// Passwords match. Generate new JWT.
				//
				token, err := generateToken(selectedUser)
				if err == nil {
					c.JSON(200, gin.H{"token": token})
				} else {
					c.JSON(401, gin.H{"error": "cannot generate token, verify if the correct key is used"})
				}
			} else {
				c.JSON(401, gin.H{"error": "username and/or password do not match"})
			}
		} else {
			c.JSON(401, gin.H{"error": "username and/or password do not match"})
		}
	} else {
		c.JSON(422, gin.H{"error": "required field(s) are empty"})
	}
}
