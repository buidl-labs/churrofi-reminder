package handlers

import (
	"net/http"
	"reminders/database"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func ReminderGet(DB *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		reminders, err := database.GetAllReminders(DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, &gin.H{
				"reminders": reminders,
			})

		}
	}
}
