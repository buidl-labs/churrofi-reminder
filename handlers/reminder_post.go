package handlers

import (
	"net/http"
	"reminders/database"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type AddReminderData struct {
	Action string `json:"action"`
	Email  string `json:"email"`
}

func ReminderPost(DB *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var postData AddReminderData
		c.Bind(&postData)
		isEmail := valid.IsEmail(postData.Email)
		if !isEmail {
			c.JSON(http.StatusBadRequest, &gin.H{
				"error": "email not valid",
			})
		} else {

			reminder, err := database.AddReminder(DB, postData.Action, postData.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, &gin.H{
					"error": "reminder not created",
				})
			} else {
				c.JSON(http.StatusCreated, &gin.H{
					"reminder": reminder,
				})

			}
		}

	}
}
