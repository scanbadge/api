package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//R is gin-gonic's engine used for controlling gin
var R *gin.Engine

//AuthRequired is middleware to protect routes from unauthorized users
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if true { //do some validation logic here
			//login is wrong, stop.
			c.Redirect(302, "/?pleaseloginfirst")
			c.Abort()
		} else {
			//login is ok, proceed
			c.Next()
		}
	}
}

func setupRouter() {
	R = gin.Default()
	R.LoadHTMLGlob("templates/*")
	R.Static("/static", "static")

	R.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", nil)
	})

	//validationGroup := R.Group("/verify", AuthRequired()) uncomment if this needs to be protected
	validationGroup := R.Group("/verify")
	{
		validationGroup.GET("/:door/:id", verifyTag)
	}

	R.Run("127.0.0.1:4700")
}

func main() {
	fmt.Println("Started!")
	setupRouter()
}

func verifyTag(c *gin.Context) {
	door := c.Param("door")
	id := c.Param("id")

	result := "Welcome!"
	c.JSON(200, gin.H{
		"result": result,
		"door":   door,
		"id":     id,
	})
}
