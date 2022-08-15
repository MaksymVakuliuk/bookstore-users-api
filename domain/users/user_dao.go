package users

import (
	"fmt"

	"github.com/MaksymVakuliuk/bookstore-users-api/datasources/mysqldb"
	"github.com/MaksymVakuliuk/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUserById  = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser   = "UPDATE users SET first_name = ?, last_name = ?, email = ?, status = ? WHERE ID = ?;"
	queryDereleUser   = "DELETE FROM users WHERE id = ?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := mysqldb.UsersDB.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
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
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
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

func (user *User) SearchUserByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := mysqldb.UsersDB.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.ParseMySQLError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
