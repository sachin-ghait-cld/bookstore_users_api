package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sachin-ghait-cld/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows"
)

// ParseError parse mysql rest_errors
func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("Error parsing database response", sqlErr)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("Invalid data")
	}
	return rest_errors.NewInternalServerError("error processing request", sqlErr)
}
