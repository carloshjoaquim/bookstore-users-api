package services

import (
	"github.com/carloshjoaquim/bookstore-users-api/domain/users"
	"github.com/carloshjoaquim/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil
}