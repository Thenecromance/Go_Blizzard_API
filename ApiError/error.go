package ApiError

import (
	"errors"
)

var (
	ErrorNotSupported   = errors.New("This feature is not supported by the CN region")
	ErrorNotImplemented = errors.New("This feature is in developing now")
	ErrorInDevelopment  = errors.New("This feature is in developing now")
)

var (
	ErrorNoPlayerToken            = errors.New("No valid token found")
	ErrorRequestFailed            = errors.New("Request to CN API failed")
	ErrorPlayerNotAllowedToAccess = errors.New("Player is not allowed to access this data")
)
