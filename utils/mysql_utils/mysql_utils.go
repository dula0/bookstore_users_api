package mysql_utils

import (
	"strings"

	"github.com/dula0/bookstore_users_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsError = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NotFoundError("no match for given id")
		}
		return errors.InternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("invalid data")
	}
	return errors.InternalServerError("error processing request")
}
