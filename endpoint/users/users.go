package users

import (
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
)

// User describes a user. Uses mapping for database and json.
type User struct {
	ID        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Firstname string `db:"Firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	IsAdmin   bool   `db:"admin" json:"admin"`
}
