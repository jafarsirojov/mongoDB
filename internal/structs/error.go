package structs

import "errors"

var (
	ErrInternal   = errors.New("Internal server error")
	ErrNotFound   = errors.New("Not found")
	ErrBadRequest = errors.New("Bad request")
)
