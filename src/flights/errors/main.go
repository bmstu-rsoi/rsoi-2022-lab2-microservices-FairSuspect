package errors

import (
	"errors"
)

var (
	RecordNotFound  = errors.New("Record is not found")
	UnknownError    = errors.New("Unknown error")
	InvalidRequest  = errors.New("Invalid request data")
	DBAdditionError = errors.New("Eror in addition new record to DB")
)
