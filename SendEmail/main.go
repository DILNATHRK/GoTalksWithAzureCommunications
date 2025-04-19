package main

import (
	"github.com/DILNATHRK/GoWithAzureCommunications/SendEmail/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/send-email", handler.SendEmailHandler)

	router.Run(":8080")
}
