package users

import (
	"github.com/carloshjoaquim/bookstore-users-api/domain/users"
	"github.com/carloshjoaquim/bookstore-users-api/services"
	"github.com/carloshjoaquim/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {

}
