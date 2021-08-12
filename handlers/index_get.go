package handlers

import "github.com/gin-gonic/gin"

func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, &gin.H{
			"message": "hey",
		})
	}
}
