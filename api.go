package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoint/devices"
	"github.com/scanbadge/api/endpoint/logs"
	"github.com/scanbadge/api/endpoint/users"
	"log"
	"strconv"
)

func main() {
	// Command-line flags.
	debug := flag.Bool("d", false, "If true, the program will output debug information.")
	flag.Parse()

	// Initialize all required configuration before starting gin.
	fmt.Println("[Scanbadge] Starting ScanBadge API...")
	configuration.Read()
	configuration.ReadKey()
	configuration.Dbmap = initDb()

	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.POST("/auth", authentication.Authenticate)

	authorized := router.Group("/", authentication.AuthRequired())
	{
		authorized.GET("/devices", devices.GetDevices)
		authorized.GET("/devices/:id", devices.GetDevice)
		authorized.POST("/devices", devices.AddDevice)
		authorized.PUT("/devices/:id", devices.UpdateDevice)
		authorized.DELETE("/devices/:id", devices.DeleteDevice)

		authorized.GET("/logs", logs.GetLogs)
		authorized.GET("/logs/:id", logs.GetLog)
		authorized.POST("/logs", logs.AddLog)
		authorized.PUT("/logs/:id", logs.UpdateLog)
		authorized.DELETE("/logs/:id", logs.DeleteLog)

		authorized.GET("/users", users.GetUsers)
		authorized.GET("/users/:id", users.GetUser)
		authorized.POST("/users", users.AddUser)
		authorized.PUT("/users/:id", users.UpdateUser)
		authorized.DELETE("/users/:id", users.DeleteUser)
	}

	// By default, gin will listen 'n serve on localhost:8080. Edit config.json to apply changes.
	host := fmt.Sprintf("%s:%s", configuration.Config.ServerHost, strconv.Itoa(configuration.Config.ServerPort))
	fmt.Println(fmt.Sprintf("[ScanBadge] Listening and serving HTTP on %s", host))
	router.Run(host)
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
	Dbmap.AddTableWithName(logs.Log{}, "logs").SetKeys(true, "ID")
	Dbmap.AddTableWithName(users.User{}, "users").SetKeys(true, "ID").SetUniqueTogether("username", "email")
	err = Dbmap.CreateTablesIfNotExists()
	checkErr("Creating table failed", err)

	return Dbmap
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
