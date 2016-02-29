package utility

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
)

// ReadData reads input from the specified file path.
func ReadData(filename string) ([]byte, error) {
	if filename == "" {
		return nil, fmt.Errorf("No filename specified")
	}

	return ioutil.ReadFile(filename)
}

// HashPassword hashes a password with the bcrypt algorithm using a cost of 10.
func HashPassword(password string) string {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
