package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/endpoints"
	"github.com/scanbadge/api/models"
	"github.com/scanbadge/api/utility"
	"gopkg.in/gorp.v1"
)

var debug, adduser bool

func init() {
	// Command-line flags.
	flag.BoolVar(&debug, "d", false, "If this flag is applied, output debug information.")
	flag.BoolVar(&adduser, "add-user", false, "If this flag is applied, add a new user.")
	flag.Parse()

	// Initialize all required configuration before starting gorp or gin.
	fmt.Println("[Scanbadge] Starting ScanBadge API...")
	configuration.Read()
	configuration.ReadKey()
	configuration.Dbmap = initDb()
}

func main() {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	if adduser {
		if !addUser() {
			return
		}
	}

	router := gin.New()

	router.POST("/auth", authentication.Authenticate)
	router.Use(CORSMiddleware())

	authorized := router.Group("/", authentication.AuthRequired())
	{
		authorized.GET("/actions", endpoints.GetActions)
		authorized.GET("/actions/:id", endpoints.GetAction)
		authorized.POST("/actions", endpoints.AddAction)
		authorized.PUT("/actions/:id", endpoints.UpdateAction)
		authorized.DELETE("/actions/:id", endpoints.DeleteAction)

		authorized.GET("/conditions", endpoints.GetConditions)
		authorized.GET("/conditions/:id", endpoints.GetCondition)
		authorized.POST("/conditions", endpoints.AddCondition)
		authorized.PUT("/conditions/:id", endpoints.UpdateCondition)
		authorized.DELETE("/conditions/:id", endpoints.DeleteCondition)

		authorized.GET("/count", endpoints.GetAllCount)
		authorized.GET("/count/:id", endpoints.GetCount)

		authorized.GET("/devices", endpoints.GetDevices)
		authorized.GET("/devices/:id", endpoints.GetDevice)
		authorized.POST("/devices", endpoints.AddDevice)
		authorized.PUT("/devices/:id", endpoints.UpdateDevice)
		authorized.DELETE("/devices/:id", endpoints.DeleteDevice)

		authorized.GET("/logs", endpoints.GetLogs)
		authorized.GET("/logs/:id", endpoints.GetLog)
		authorized.POST("/logs", endpoints.AddLog)
		authorized.PUT("/logs/:id", endpoints.UpdateLog)
		authorized.DELETE("/logs/:id", endpoints.DeleteLog)

		authorized.GET("/users", endpoints.GetUsers)
		authorized.GET("/users/:id", endpoints.GetUser)
		authorized.POST("/users", endpoints.AddUser)
		authorized.PUT("/users/:id", endpoints.UpdateUser)
		authorized.DELETE("/users/:id", endpoints.DeleteUser)
	}

	// By default, gin will listen 'n serve on localhost:8080. Edit config.json to apply changes.
	host := fmt.Sprintf("%s:%s", configuration.Config.ServerHost, strconv.Itoa(configuration.Config.ServerPort))
	fmt.Println(fmt.Sprintf("[ScanBadge] Listening and serving HTTP on %s", host))
	router.Run(host)
}

func initDb() *gorp.DbMap {
	// Must be in the following format: username:password@protocol(address:port)/dbname?param1=value1&...&paramN=valueN
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=true",
		configuration.Config.Database.Username,
		configuration.Config.Database.Password,
		configuration.Config.Database.Protocol,
		configuration.Config.Database.Host,
		configuration.Config.Database.Port,
		configuration.Config.Database.DatabaseName,
		configuration.Config.Database.Charset)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	// Firewall might block database connection, but it will not be shown, so ensure that we receive an error.
	if port, err := strconv.Atoi(configuration.Config.Database.Port); err == nil {
		if !utility.IsPortOpen("tcp", configuration.Config.Database.Host, port) {
			panic("Cannot connect to MySQL database...")
		}
	} else {
		panic(err)
	}

	Dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: configuration.Config.Database.Engine, Encoding: configuration.Config.Database.Encoding}}

	Dbmap.AddTableWithName(models.Action{}, "actions").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.ActionType{}, "action_types").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.Condition{}, "conditions").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.ConditionType{}, "condition_types").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.Device{}, "devices").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.Log{}, "logs").SetKeys(true, "ID")
	Dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "ID").SetUniqueTogether("user_username", "user_email")

	if err = Dbmap.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	return Dbmap
}

// CORSMiddleware is middleware that sets several Access-Control headers to allow increased security.
// The following headers are set (with default values):
// Access-Control-Allow-Origin: * OR <config.json `AllowedOrigins` value>
// Access-Control-Max-Age: 86400
// Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE, UPDATE
// Access-Control-Allow-Headers: Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
// Access-Control-Expose-Headers: Content-Length

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := configuration.Config.AllowedOrigins

		if debug || origin == "" {
			origin = "*"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
