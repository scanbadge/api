package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/users"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// AuthRequired is middleware that validates whether or not the current request has been authenticated by a JWT.
// If the request is not authenticated, a 401 HTTP code will be returned.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header and trim ' Bearer '.
		var encodedToken = c.Request.Header.Get("Authorization")

		parsedToken, err := ParseToken(encodedToken)
		if err == nil {
			if parsedToken.Valid {
				// Authentication is OK, proceed.
				c.Next()
				return
			}
		}

		// Authentication is wrong, stop.
		log.Println(err)
		c.JSON(401, gin.H{"error": "Unauthorized: Access is denied"})
		c.Abort()
	}
}

// Authenticate authenticates a user by providing a JWT if the provided email and password match with the information from the database.
func Authenticate(c *gin.Context) {
	var user users.User
	err := c.BindJSON(&user)

	if err == nil && user.Username != "" && user.Password != "" && len(user.Password) < 256 {
		// TODO: add brute force protection
		var selectedUser users.User
		err := configuration.Dbmap.SelectOne(&selectedUser, "select * from users where username = ?", user.Username)

		if err == nil && selectedUser.Password != "" {
			err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))

			if err == nil {
				//
				// Passwords match. Generate new JWT.
				//
				token, err := generateToken(selectedUser)
				if token != "" && err == nil {
					c.JSON(200, gin.H{"token": token})
				} else {
					c.JSON(401, gin.H{"error": "cannot generate a token"})
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