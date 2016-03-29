package endpoints

import "github.com/gin-gonic/gin"

func showError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
	c.Abort()
}

func showSucces(c *gin.Context, message string) {
	c.JSON(200, gin.H{"success": message})
	c.Abort()
}

func showResult(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
	c.Abort()
}
