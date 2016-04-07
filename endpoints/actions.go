package endpoints

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/authentication"
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
)

// GetActions gets all actions.
func GetActions(c *gin.Context) {
	var actions []models.Action
	uid, err := authentication.GetUserID(c)

	if err == nil {
		// select actions.* from actions
		// inner join action_types on actions.action_type_id=action_types.action_type_id
		// where actions.user_id=?
		_, err := configuration.Dbmap.Select(&actions, "select actions.* from actions inner join action_types on actions.action_type_id=action_types.action_type_id where actions.user_id=?", uid)

		if err == nil && len(actions) > 0 {
			showResult(c, 200, actions)
			return
		}
	}

	showError(c, 404, fmt.Errorf("actions not found"))
}

// GetAction gets an action based on the provided identifier.
func GetAction(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var action models.Action
		uid, err := authentication.GetUserID(c)

		if err == nil {
			err := configuration.Dbmap.SelectOne(&action, "select * from actions where id=? and user_id=?", id, uid)

			if err == nil {
				showResult(c, 200, action)
				return
			}
		}

		showError(c, 404, fmt.Errorf("action not found"))
		return
	}

	showError(c, 422, fmt.Errorf("action identifier not provided"))
}

// AddAction adds a new action.
func AddAction(c *gin.Context) {
	var action models.Action
	err := c.Bind(&action)

	if err == nil {
		if action.Name != "" && action.Description != "" {
			// Verify if provided ActionTypeID exists.
			count, err := configuration.Dbmap.SelectInt("select count(*) from action_types where id=?", action.ActionTypeID)

			if count == 1 && err == nil {
				uid, err := authentication.GetUserID(c)

				if err == nil {
					action.UserID = uid
					err := configuration.Dbmap.Insert(&action)

					if err == nil {
						showResult(c, 201, action)
						return
					}
				}
			}

			showError(c, 400, fmt.Errorf("adding new action failed"))
			return
		}

		showError(c, 422, fmt.Errorf("field(s) are empty"))
		return
	}

	showError(c, 400, fmt.Errorf("adding new action failed"))
}

// UpdateAction updates an action based on the identifer.
func UpdateAction(c *gin.Context) {

}

// DeleteAction deletes an action based on the identifier.
func DeleteAction(c *gin.Context) {

}
