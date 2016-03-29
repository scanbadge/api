package endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
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
	var logs []models.Log
	uid, err := authentication.GetUserID(c)

	if err == nil {
		_, err := configuration.Dbmap.Select(&logs, "select * from logs where user_id=?", uid)

		if err == nil && len(logs) > 0 {
			showResult(c, 200, logs)
		}
	}

	showError(c, 404, fmt.Errorf("no log entries found"))
}

// GetLog gets a device based on the provided identifier.
func GetLog(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var log models.Log
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
	var log models.Log
	err := c.Bind(&log)

	if err == nil {
		if log.Type != "" && log.Description != "" && log.Origin != "" && log.Object != "" {
			uid, err := authentication.GetUserID(c)

			if err == nil {
				log.UserID = uid
				err := configuration.Dbmap.Insert(&log)

				if err == nil {
					showResult(c, 201, log)
				}
			}

			showError(c, 400, fmt.Errorf("adding new log entry failed"))
		}

		showError(c, 422, fmt.Errorf("field(s) are empty"))
	}

	showError(c, 400, fmt.Errorf("adding new log entry failed"))
}

// UpdateLog updates a device based on the identifer.
func UpdateLog(c *gin.Context) {
	c.JSON(403, gin.H{"error": "PUT /logs is not supported yet"})
}

// DeleteLog deletes a device based on the identifier.
func DeleteLog(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var log models.Log
		err := c.BindJSON(&log)

		if err == nil {
			uid, err := authentication.GetUserID(c)

			if err == nil {
				log.UserID = uid
				count, err := configuration.Dbmap.Delete(&log)

				if err == nil && count == 1 {
					showSucces(c, fmt.Sprintf("log entry with id %s is deleted", id))
				}
			}
		}

		showError(c, 400, fmt.Errorf("deleting log entry failed"))
	}

	showError(c, 400, fmt.Errorf("deleting log entry failed due to missing identifier"))
}
