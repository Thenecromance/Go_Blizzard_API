package ApiError

import (
	"errors"
)

var (
	ErrorNotSupported   = errors.New("This feature is not supported by the CN region")
	ErrorNotImplemented = errors.New("This feature is in developing now")
)
