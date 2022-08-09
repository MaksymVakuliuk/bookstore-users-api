package services

import (
	"github.com/MaksymVakuliuk/bookstore-users-api/domain/users"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userid int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
