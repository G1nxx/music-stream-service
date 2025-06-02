package storage

import (
	"errors"
)

var (
	ErrTrackNotFound = errors.New("track not found")
	ErrUserNotFound  = errors.New("user not found")
)
