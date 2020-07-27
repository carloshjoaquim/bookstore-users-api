package app

import (
	"github.com/carloshjoaquim/bookstore-users-api/controllers/ping"
	"github.com/carloshjoaquim/bookstore-users-api/controllers/users"
)

func mapUrl() {

	// Ping
	router.GET("/ping", ping.Ping)

	// Users
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)

}
