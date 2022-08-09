package users

import (
	"fmt"

	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user &d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	if userDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user $d alredy exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
