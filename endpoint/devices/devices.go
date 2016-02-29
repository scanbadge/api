package devices

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/configuration"
)

// Device describes a device. Uses mapping for database and json.
type Device struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Key  string `db:"key" json:"key"`
}

// GetDevices gets all devices.
func GetDevices(c *gin.Context) {
	var devices []Device
	_, err := configuration.Dbmap.Select(&devices, "select * from devices")

	if err == nil {
		c.JSON(200, devices)
	} else {
		c.JSON(404, gin.H{"error": "no device(s) found"})
	}
}

// GetDevice gets a device based on the provided identifier.
func GetDevice(c *gin.Context) {
	id := c.Params.ByName("id")
	var device Device
	err := configuration.Dbmap.SelectOne(&device, "select * from devices where id=?", id)

	if err == nil {
		c.JSON(200, device)
	} else {
		c.JSON(404, gin.H{"error": "device not found"})
	}
}

// AddDevice adds a new device.
func AddDevice(c *gin.Context) {
	var device Device
	err := c.BindJSON(&device)

	if err == nil {
		if device.Key != "" && device.Name != "" {
			err := configuration.Dbmap.Insert(&device)

			if err == nil {
				c.JSON(201, device)
			} else {
				// err.Error() should be removed as soon as we have implemented
				// better error handling when adding resources to API.
				c.JSON(400, gin.H{"error": "Adding new device failed due to " + err.Error()})
			}
		} else {
			c.JSON(422, gin.H{"error": "field(s) are empty"})
		}
	} else {
		// err.Error() should be removed as soon as we have implemented
		// better error handling when adding resources to API.
		c.JSON(400, gin.H{"error": "Adding new device failed due to " + err.Error()})
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
