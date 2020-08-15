package app

import (
	"github.com/carloshjoaquim/bookstore-users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default() // Creates a new goroutine for each request.
)

func StartApplication() {
	mapUrl()
	logger.Info("Start the application...")
	router.Run(":8081")
}
