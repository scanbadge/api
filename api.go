package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/devices"
	"github.com/scanbadge/api/endpoint/users"
	"log"
	"strconv"
)

func main() {
	configuration.Read()
	configuration.Dbmap = initDb()
	router := gin.Default()

	v1 := router.Group("api/v1", AuthRequired())
	{
		// Authentication
		v1.POST("/auth", Authenticate)
		// Devices
		v1.GET("/devices", devices.GetDevices)
		v1.GET("/devices/:id", devices.GetDevice)
		v1.POST("/devices", devices.AddDevice)
		v1.PUT("/devices/:id", devices.UpdateDevice)
		v1.DELETE("/devices/:id", devices.DeleteDevice)
	}

	// By default, gin will listen 'n serve on localhost:8080. Edit config.json to apply changes.
	router.Run(fmt.Sprintf("%s:%s", configuration.Config.ServerHost, strconv.Itoa(configuration.Config.ServerPort)))
}

func initDb() *gorp.DbMap {
	// Must be in the following format: username:password@protocol(address)/dbname?param1=value1&...&paramN=valueN
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=true",
		configuration.Config.Database.Username,
		configuration.Config.Database.Password,
		configuration.Config.Database.Protocol,
		configuration.Config.Database.Host,
		configuration.Config.Database.Port,
		configuration.Config.Database.DatabaseName,
		configuration.Config.Database.Charset)

	db, err := sql.Open("mysql", dsn)

	checkErr("Cannot open connection to database", err)

	Dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: configuration.Config.Database.Engine, Encoding: configuration.Config.Database.Encoding}}

	Dbmap.AddTableWithName(devices.Device{}, "devices").SetKeys(true, "ID")
	Dbmap.AddTableWithName(users.User{}, "users").SetKeys(true, "ID").SetUniqueTogether("Username", "Email")
	err = Dbmap.CreateTablesIfNotExists()
	checkErr("Creating table failed", err)

	return Dbmap
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
