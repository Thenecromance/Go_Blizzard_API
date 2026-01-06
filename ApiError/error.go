package ApiError

import (
	"errors"
)

var (
	ErrorNotSupported   = errors.New("This feature is not supported by the CN region")
	ErrorNotImplemented = errors.New("This feature is in developing now")
	ErrorInDevelopment  = errors.New("This feature is in developing now")
	ErrorNoToken        = errors.New("No valid token found")
)
