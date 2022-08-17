package services

import (
	"github.com/MaksymVakuliuk/bookstore-users-api/domain/users"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/crypto"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/date"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

var (
	UserService userServiceInteface = &userService{}
)

type userService struct{}

type userServiceInteface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	GetUser(userid int64) (*users.User, *errors.RestErr)
	UpdateUser(isPartinal bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	SearchUserByStatus(status string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userid int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUser(isPartinal bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
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
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) SearchUserByStatus(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.SearchUserByStatus(status)
}
