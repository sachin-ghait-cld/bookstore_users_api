package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sachin-ghait-cld/bookstore_users_api/utils/errors"
)

const (
	ErrorNoRows = "no rows"
)

// ParseError parse mysql errors
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
