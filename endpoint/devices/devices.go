package devices

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
)

// Device describes a device. Uses mapping for database and json.
type Device struct {
	ID     int64  `db:"id" json:"id"`
	UserID int64  `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
	Key    string `db:"key" json:"key"`
}

type addDevice struct {
	Name string `json:"name" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

// GetDevices gets all devices.
func GetDevices(c *gin.Context) {
	var devices []Device
	userID, err := authentication.GetUserID(c)

	if err == nil {
		_, err := configuration.Dbmap.Select(&devices, "select * from devices where user_id=?", userID)

		if err == nil && len(devices) > 0 {
			c.JSON(200, devices)
		} else {
			c.JSON(404, gin.H{"error": "device(s) not found"})
		}
	} else {
		c.JSON(404, gin.H{"error": "device(s) not found"})
	}
}

// GetDevice gets a device based on the provided identifier.
func GetDevice(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var device Device
		userID, err := authentication.GetUserID(c)

		if err == nil {
			err := configuration.Dbmap.SelectOne(&device, "select * from devices where id=? and user_id=?", id, userID)

			if err == nil {
				c.JSON(200, device)
			} else {
				c.JSON(404, gin.H{"error": "device not found"})
			}
		} else {
			c.JSON(404, gin.H{"error": "device not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "no identifier provided"})
	}
}

// AddDevice adds a new device.
func AddDevice(c *gin.Context) {
	var (
		device    Device
		addDevice addDevice
	)

	err := c.BindJSON(&addDevice)

	if err == nil {
		if addDevice.Key != "" && addDevice.Name != "" {
			device.Key = addDevice.Key
			device.Name = addDevice.Name

			userID, err := authentication.GetUserID(c)

			if err == nil {
				device.UserID = userID

				err := configuration.Dbmap.Insert(&device)

				if err == nil {
					c.JSON(201, device)
				} else {
					c.JSON(400, gin.H{"error": "adding new device failed due to " + err.Error()})
				}
			} else {
				c.JSON(400, gin.H{"error": "adding new device failed due to " + err.Error()})
			}
		} else {
			c.JSON(422, gin.H{"error": "field(s) are empty"})
		}
	} else {
		c.JSON(400, gin.H{"error": "adding new device failed due to " + err.Error()})
	}
}

// UpdateDevice updates a device based on the identifer.
func UpdateDevice(c *gin.Context) {
	c.JSON(403, gin.H{"error": "PUT is not supported"})
}

// DeleteDevice deletes a device based on the identifier.
func DeleteDevice(c *gin.Context) {
	c.JSON(403, gin.H{"error": "DELETE is not supported"})
}
