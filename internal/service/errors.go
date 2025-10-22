package service

import "errors"

var (
	ErrUserExists      = errors.New("user already registered")
	ErrInvalidPassword = errors.New("invalid password")
	ErrServerError     = errors.New("internal server error")
)
