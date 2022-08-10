package users

import (
	"fmt"
	"strings"

	"github.com/MaksymVakuliuk/bookstore-users-api/datasources/mysqldb"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/date"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

const (
	indexUniqueEmail          = "email_UNIQUE"
	internalServerErrorFormat = "error when trying to save user: %s"
	errorNoRows               = "no rows in result set"
	queryInsertUser           = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	queryGetUserById          = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user id = %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user by id = %d : %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exicts", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf(internalServerErrorFormat, err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf(internalServerErrorFormat, err.Error()))
	}
	user.Id = userId
	return nil
}
