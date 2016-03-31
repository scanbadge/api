package endpoints

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
)

// GetDevices gets all devices.
func GetDevices(c *gin.Context) {
	var devices []models.Device
	userID, err := authentication.GetUserID(c)

	if err == nil {
		_, err := configuration.Dbmap.Select(&devices, "select devices.* from devices inner join users on devices.user_id=users.id where devices.user_id=?", userID)

		if err == nil && len(devices) > 0 {
			showResult(c, 200, devices)
			return
		}
	}

	log.Println(err)
	showError(c, 404, fmt.Errorf("device(s) not found"))
}

// GetDevice gets a device based on the provided identifier.
func GetDevice(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var device models.Device
		userID, err := authentication.GetUserID(c)

		if err == nil {
			err := configuration.Dbmap.SelectOne(&device, "select * from devices where id=? and user_id=?", id, userID)

			if err == nil {
				showResult(c, 200, device)
				return
			}
		}

		log.Println(err)
		showError(c, 404, fmt.Errorf("device not found"))
		return
	}

	showError(c, 422, fmt.Errorf("no identifier provided"))
}

// AddDevice adds a new device.
func AddDevice(c *gin.Context) {
	var device models.Device

	err := c.Bind(&device)

	if err == nil && device.Description != "" && device.Key != "" && device.Name != "" {
		userID, err := authentication.GetUserID(c)

		if err == nil {
			device.UserID = userID

			err := configuration.Dbmap.Insert(&device)

			if err == nil {
				showResult(c, 201, device)
				return
			}
		}

		showError(c, 400, fmt.Errorf("adding new device failed"))
		return
	}

	showError(c, 422, fmt.Errorf("field(s) are empty"))
}

// UpdateDevice updates a device based on the identifer.
func UpdateDevice(c *gin.Context) {
	showError(c, 403, fmt.Errorf("PUT /devices is not supported"))
}

// DeleteDevice deletes a device based on the identifier.
func DeleteDevice(c *gin.Context) {
	showError(c, 403, fmt.Errorf("DELETE /devices is not supported"))
}
