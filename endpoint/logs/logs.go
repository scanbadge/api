package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
)

// Log describes a log entry. Uses mapping for database and json.
type Log struct {
	ID      int64  `db:"id" json:"id"`
	UserID  int64  `db:"user_id" json:"user_id"`
	Message string `db:"message" json:"message"`
	Origin  string `db:"origin" json:"origin"`
}

// GetLogs gets all devices.
func GetLogs(c *gin.Context) {
	var logs []Log
	uid, err := authentication.GetUserID(c)

	if err == nil {
		_, err := configuration.Dbmap.Select(&logs, "select * from logs where user_id=?", uid)

		if err == nil && len(logs) > 0 {
			c.JSON(200, logs)
		} else {
			c.JSON(404, gin.H{"error": "no log entries found"})
		}
	} else {
		c.JSON(404, gin.H{"error": "no log entries found"})
	}
}

// GetLog gets a device based on the provided identifier.
func GetLog(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var log Log
		uid, err := authentication.GetUserID(c)

		if err == nil {
			err := configuration.Dbmap.SelectOne(&log, "select * from logs where id=? and user_id=?", id, uid)

			if err == nil {
				c.JSON(200, log)
			} else {
				c.JSON(404, gin.H{"error": "log entry not found"})
			}
		} else {
			c.JSON(404, gin.H{"error": "log entry not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "no log entry identifier provided"})
	}
}

// AddLog adds a new log entry for the current user.
func AddLog(c *gin.Context) {
	var log Log
	err := c.BindJSON(&log)

	if err == nil {
		if log.Message != "" && log.Origin != "" {
			uid, err := authentication.GetUserID(c)

			if err == nil {
				log.UserID = uid
				err := configuration.Dbmap.Insert(&log)

				if err == nil {
					c.JSON(201, log)
				} else {
					c.JSON(400, gin.H{"error": "adding new log entry failed due to " + err.Error()})
				}
			} else {
				c.JSON(400, gin.H{"error": "adding new log entry failed"})
			}
		} else {
			c.JSON(422, gin.H{"error": "field(s) are empty"})
		}
	} else {
		// err.Error() should be removed as soon as we have implemented
		// better error handling when adding resources to API.
		c.JSON(400, gin.H{"error": fmt.Sprintf("adding new log entry failed due to %s", err.Error())})
	}
}

// UpdateLog updates a device based on the identifer.
func UpdateLog(c *gin.Context) {
	c.JSON(403, gin.H{"error": "PUT /logs for is not supported yet"})
}

// DeleteLog deletes a device based on the identifier.
func DeleteLog(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var log Log
		err := c.BindJSON(&log)

		if err == nil {
			uid, err := authentication.GetUserID(c)

			if err == nil {
				log.UserID = uid
				count, err := configuration.Dbmap.Delete(&log)

				if err == nil && count == 1 {
					c.JSON(200, gin.H{"success": fmt.Sprintf("log entry with id %s is deleted", id)})
				} else {
					c.JSON(400, gin.H{"error": "deleting log entry failed"})
				}
			} else {
				c.JSON(400, gin.H{"error": fmt.Sprintf("deleting log entry failed due to %s", err.Error())})
			}
		} else {
			c.JSON(400, gin.H{"error": fmt.Sprintf("deleting log entry failed due to %s", err.Error())})
		}
	} else {
		c.JSON(400, gin.H{"error": "deleting log entry failed due to missing identifier"})
	}
}
