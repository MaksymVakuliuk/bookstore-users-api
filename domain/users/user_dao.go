package users

import (
	"github.com/MaksymVakuliuk/bookstore-users-api/datasources/mysqldb"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/date"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser  = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE ID = ?;"
	queryDereleUser  = "DELETE FROM users WHERE id = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.ParseMySQLError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date.GetNowDBFormat()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.ParseMySQLError(err)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.ParseMySQLError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.ParseMySQLError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryDereleUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return errors.ParseMySQLError(err)
	}
	return nil
}
