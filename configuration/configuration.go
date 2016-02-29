package configuration

import (
	"encoding/json"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql" // implement MySQL SQL driver
	"io/ioutil"
)

// Config returns the current configuration settings.
var Config Configuration

// Dbmap contains a pointer to the gorp.DpMap
var Dbmap *gorp.DbMap

// Configuration stores the main configuration for the application.
type Configuration struct {
	ServerHost string // The hostname on which the HTTP server will run, e.g. 'localhost'.
	ServerPort int    // The port on which the HTTP server will run, e.g. '8080'.
	Key        string // The relative path of the hex-encoded key used for signing JWT. The key must be at least 256 bits in length.
	Database   MySQLConfiguration
}

// MySQLConfiguration stores the specific MySQL configuration for this application.
type MySQLConfiguration struct {
	Username     string // The username of the MySQL user.
	Password     string // The password of the MySQL user.
	DatabaseName string // The MySQL database to use.
	Protocol     string // The protocol to use to connect to the MySQL server, either 'tcp' or 'udp'.
	Host         string // The host to use to connect to the MySQL server, e.g. 'localhost'.
	Port         string // The port to use to connect to the MySQL server, e.g. '3306'.
	Charset      string // The charset to use for connection string, e.g. 'utf8mb4,utf8'.
	Engine       string // The database engine of the MySQL server, e.g. 'InnoDB'.
	Encoding     string // The encoding used on the MySQL server, e.g. 'UTF8'.
}

// Read reads the configuration file and parses it. If any errors are found, a panic will occur.
func Read() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err.Error())
	}
}
