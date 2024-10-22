package box

import "github.com/pkg/errors"

var errorsMap map[error]int = make(map[error]int)

func NewError(msg string, status int) error {
	newErr := errors.New(msg)
	errorsMap[newErr] = status
	return newErr
}

func GetStatus(err error) int {
	causeErr := errors.Cause(err)
	return errorsMap[causeErr]
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}
