package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() // Creates a new goroutine for each request.
)

func StartApplication() {
	mapUrl()
	router.Run(":8080")
}
