package types

import (
	"fmt"
	"net/http"
)

type ErrorString string

func (s ErrorString) Error() string {
	return string(s)
}

const (
	ErrNotExist ErrorString = "does not exist"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, msgFormat string, args ...interface{}) *Error {
	message := fmt.Sprintf(msgFormat, args...)
	if len(message) == 0 {
		message = http.StatusText(code % 1000)
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}
