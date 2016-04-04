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

// GetAllCount gets all the count of objects of current user for all endpoints.
func GetAllCount(c *gin.Context) {
	uid, err := authentication.GetUserID(c)
	endpoints := []string{"actions", "action_error", "action_success", "conditions", "devices", "logs", "users"}
	var counts []models.Count

	if err == nil {
		for _, endpoint := range endpoints {
			var count int64

			if endpoint == "action_error" || endpoint == "action_success" {
				count, err = configuration.Dbmap.SelectInt(fmt.Sprintf("select count(log_type) from logs where log_type='%s' and user_id=?", endpoint), uid)
			} else {
				count, err = configuration.Dbmap.SelectInt(fmt.Sprintf("select count(*) from %s where user_id=?", endpoint), uid)
			}

			if err == nil {
				counts = append(counts, models.Count{Count: count, Endpoint: endpoint})
			} else {
				log.Println(err)
			}
		}

		showResult(c, 200, counts)
		return
	}

	showError(c, 404, fmt.Errorf("no endpoints found to count"))
}

// GetCount gets the count of objects of current user for the provided endpoint.
// For example, when a user has 10 actions, the count will be 10.
// An error will be returned if the endpoint is not found.
func GetCount(c *gin.Context) {
	id := c.Params.ByName("id")
	endpoints := []string{"actions", "conditions", "devices", "logs", "users"}

	if containsString(endpoints, id) {
		uid, err := authentication.GetUserID(c)

		if err == nil {
			query := fmt.Sprintf("select count(*) from %s where user_id=?", id)
			i, err := configuration.Dbmap.SelectInt(query, uid)

			if err == nil {
				var count models.Count
				count.Count = i
				count.Endpoint = id

				showResult(c, 200, count)
				return
			}
		}

		log.Println(err)
	}

	showError(c, 404, fmt.Errorf(fmt.Sprintf("endpoint '%s' not found", id)))
}

func containsString(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}

	return false
}
