package errors

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseMySQLError(err error) *RestErr {
	mySQLErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return NewInternalServerError("no user mathching given id")
		}
		return NewInternalServerError("error parsing database response")
	}
	switch mySQLErr.Number {
	case 1062:
		return NewBadRequestError("invalid data")
	}
	return NewInternalServerError("error processing request")
}
