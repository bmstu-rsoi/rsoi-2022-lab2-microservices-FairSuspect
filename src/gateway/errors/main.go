package errors

import (
	"errors"
)

var (
	ForbiddenTicket = errors.New("Forbidden ticket for this user")
)
