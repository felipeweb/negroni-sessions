package sessions

import (
	"errors"
)

var (
	// ErrInvalidID error
	ErrInvalidID = errors.New("session: invalid session id")
	// ErrInvalidModified error
	ErrInvalidModified = errors.New("mongostore: invalid modified value")
)
