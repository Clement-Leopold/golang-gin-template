package common

import (
	"fmt"
)

// error wrapper
type BusinessError struct {
	Err     error
	Message ErrorMessage
	Code    ErrorCode
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("error: %s + %s", e.Message, e.Code)
}

type ErrorMessage string
type ErrorCode string

const (
	Database ErrorMessage = "d-s error"
	Param    ErrorMessage = "p-s error"
)
const (
	DatabaseCode ErrorCode = "51000"
	ParamCode    ErrorCode = "41000"
	Success      ErrorCode = "10000"
)

func DatabaseError(err error) *BusinessError {
	return &BusinessError{
		Err:     err,
		Code:    DatabaseCode,
		Message: Database,
	}
}

func ParamError(err error) *BusinessError {
	return &BusinessError{
		Err:     err,
		Code:    ParamCode,
		Message: Param,
	}
}
