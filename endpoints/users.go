package endpoints

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
	"github.com/scanbadge/api/utility"
)

// GetUsers gets all users.
func GetUsers(c *gin.Context) {
	var users []models.User
	_, err := configuration.Dbmap.Select(&users, "select * from users")

	if err == nil && len(users) > 0 {
		// Omit password.
		for i := range users {
			users[i].Password = ""

			users[i].Roles = models.Role{ID: 1, Name: "Admin", Description: "A user with the highest permissions.", Level: 9001}
		}

		showResult(c, 200, users)
		return
	}

	log.Println(err)
	showError(c, 404, fmt.Errorf("user(s) not found"))
}

// GetUser gets a user based on the provided identifier.
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var user models.User
		err := configuration.Dbmap.SelectOne(&user, "select * from users where user_id=?", id)

		if err == nil {
			// Omit password.
			user.Password = ""
			showResult(c, 200, user)
			return
		}

		log.Println(err)
		showError(c, 404, fmt.Errorf("user not found"))
		return
	}

	showError(c, 422, fmt.Errorf("no identifier provided"))
}

// AddUser adds a new user.
func AddUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err == nil && user.Username != "" && user.Email != "" && user.Password != "" && user.FirstName != "" && user.LastName != "" {
		if len(user.Password) >= 12 {
			// Always hash the password when saving to the database.
			user.Password = utility.HashPassword(user.Password)

			err := configuration.Dbmap.Insert(&user)

			if err == nil {
				showResult(c, 201, user)
				return
			}
		}

		log.Println(err)
		showError(c, 400, fmt.Errorf("adding new user failed"))
		return
	}

	showError(c, 422, fmt.Errorf("field(s) are empty, adding new user failed"))
}

// UpdateUser updates a user based on the identifer.
func UpdateUser(c *gin.Context) {
	showError(c, 403, fmt.Errorf("PUT /users is not supported"))
}

// DeleteUser deletes a user based on the identifier.
func DeleteUser(c *gin.Context) {
	showError(c, 403, fmt.Errorf("DELETE /users is not supported"))
}
