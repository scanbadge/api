package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/devices"
	"log"
	"strconv"
)

func main() {
	configuration.Read()
	configuration.Dbmap = initDb()
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		// Devices
		v1.GET("/devices", devices.GetDevices)
		v1.GET("/devices/:id", devices.GetDevice)
		v1.POST("/devices", devices.AddDevice)
		v1.PUT("/devices/:id", devices.UpdateDevice)
		v1.DELETE("/devices/:id", devices.DeleteDevice)
	}

	// By default, gin will listen 'n serve on localhost:8080. Edit config.json to apply changes.
	router.Run(fmt.Sprintf("%s:%s", configuration.Config.ServerAddress, strconv.Itoa(configuration.Config.ServerPort)))
}

func initDb() *gorp.DbMap {
	// Must be in the following format: username:password@protocol(address)/dbname?param1=value1&...&paramN=valueN
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=true",
		configuration.Config.Database.Username,
		configuration.Config.Database.Password,
		configuration.Config.Database.Protocol,
		configuration.Config.Database.Host,
		configuration.Config.Database.Port,
		configuration.Config.Database.DatabaseName)

	db, err := sql.Open("mysql", dsn)

	checkErr(err, "Cannot open connection to database")

	Dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: configuration.Config.Database.Engine, Encoding: configuration.Config.Database.Encoding}}

	Dbmap.AddTableWithName(devices.Device{}, "devices").SetKeys(true, "ID")
	err = Dbmap.CreateTablesIfNotExists()
	checkErr(err, "Creating table failed")

	return Dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
