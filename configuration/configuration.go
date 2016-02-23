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
	ServerAddress string
	ServerPort    int
	Database      MySQLConfiguration
}

// MySQLConfiguration stores the specific MySQL configuration for this application.
type MySQLConfiguration struct {
	Username     string
	Password     string
	DatabaseName string
	Protocol     string
	Host         string
	Port         string
	Engine       string
	Encoding     string
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
