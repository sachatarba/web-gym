package service

import (
	"errors"
)

var (
	// Repo errors
	ErrDeleteNoSuchRow  = errors.New("can't delete, no such row for the specified identifier")
	ErrUpdateNoSuchRow  = errors.New("can't update, no such row for update for the specified identifier")
	ErrGetByIDNoSuchRow = errors.New("can't get, row for the specified identifier")
	ErrListByIDNoRows   = errors.New("can't get list, no rows for the specified identifier")
	ErrInternalDB       = errors.New("internal db error")
	ErrNoSuchClient     = errors.New("no such client for the specified login")
	ErrInternalSessionRepo = errors.New("internal session repo error")

	// Service errors
	ErrValidation = errors.New("validation error")

	// Autorization errors
	ErrWrongLogin    = errors.New("wrong login")
	ErrWrongPassword = errors.New("wrong password")
)
