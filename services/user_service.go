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

func UpdateUser(isPartinal bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartinal {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
