package devices

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// Device describes a device. Uses mapping for database and json.
type Device struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Key  string `db:"key" json:"key"`
}

// GetDevices gets all devices.
func GetDevices(c *gin.Context) {
	type Devices []Device

	var devices = Devices{
		Device{ID: 1, Name: "Arduino 1", Key: "1234"},
		Device{ID: 2, Name: "Arduino 2", Key: "1234"},
	}

	c.JSON(200, devices)
}

// GetDevice gets a device based on the identifier.
func GetDevice(c *gin.Context) {
	id := c.Params.ByName("id")

	userID, _ := strconv.ParseInt(id, 0, 64)

	switch userID {
	case 1:
		{
			content := gin.H{"id": userID, "name": "Arduino 1", "key": "1234"}
			c.JSON(200, content)
		}
	case 2:
		{
			content := gin.H{"id": userID, "name": "Arduino 2", "key": "1234"}
			c.JSON(200, content)
		}
	default:
		{
			content := gin.H{"error": "Device not found"}
			c.JSON(404, content)
		}
	}
}

// AddDevice adds a new device.
func AddDevice(c *gin.Context) {
	// To be implemented.
}

// UpdateDevice updates a device based on the identifer.
func UpdateDevice(c *gin.Context) {
	// To be implemented.
}

// DeleteDevice deletes a device based on the identifier.
func DeleteDevice(c *gin.Context) {
	// To be implemented.
}
