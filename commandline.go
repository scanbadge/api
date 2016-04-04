package main

import (
	"fmt"

	"github.com/mewbak/gopass"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
	"github.com/scanbadge/api/utility"
)

// addUser adds a new user based on the provided stdin. The provided password will be hashed using bcrypt.
// The following details are required: username, password, email, first name, last name.
func addUser() bool {
	var user models.User

	fmt.Println("Adding new user to ScanBadge API...")

	fmt.Println("Username:")
	fmt.Scanln(&user.Username)
	if user.Username == "" {
		fmt.Println("Username is required")
		return false
	}

	password, err := gopass.GetPass("Password:\n")
	if err != nil {
		fmt.Println(err)
		return false
	}

	if password == "" || len(password) < 8 || len(password) > 512 {
		fmt.Println("Password must be at least 8 characters long and cannot exceed 512 characters")
		return false
	}

	user.Password = utility.HashPassword(password)

	fmt.Println("Email:")
	fmt.Scanln(&user.Email)
	if user.Email == "" {
		fmt.Println("Email is required")
		return false
	}

	fmt.Println("First name:")
	fmt.Scanln(&user.FirstName)
	if user.FirstName == "" {
		fmt.Println("First name is required")
		return false
	}

	fmt.Println("Last name:")
	fmt.Scanln(&user.LastName)
	if user.LastName == "" {
		fmt.Println("Last name is required")
		return false
	}

	// Everything seems to be all right, attempt to insert new user to database.
	err = configuration.Dbmap.Insert(&user)

	if err == nil {
		fmt.Println(fmt.Sprintf("Successfully added new user '%s'", user.Username))
	} else {
		fmt.Println("Cannot add new user due to " + err.Error())
	}

	return err == nil
}
