package main

import (
	"net/http"

	"github.com/DILNATHRK/GoWithAzureCommunications/SendEmail/handler"
	"github.com/DILNATHRK/GoWithAzureCommunications/SendEmail/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/send-email", func(c *gin.Context) {
		var email model.EmailNotification

		if err := c.ShouldBindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "details": err.Error()})
			return
		}

		if err := handler.SendEmailWithAzureAPI(email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send email", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})
	})

	router.Run(":8080")
}
