package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/pkg/validator"
)

type Client struct {
	ID                uuid.UUID
	Login             string
	Password          string
	Fullname          string
	Email             string
	Phone             string
	Birthdate         string
	ClientMemberships []ClientMembership
	Schedules         []Schedule
}

func (cl *Client) Validate() bool {
	date, err := time.Parse(time.DateOnly, cl.Birthdate)
	if err != nil {
		return false
	}
	cl.Birthdate = date.String()

	if cl.Login == "" || cl.Password == "" {
		return false
	}

	if !validator.IsValidEmail(cl.Email) {
		return false
	}

	return validator.IsValidPhoneNumber(cl.Phone)
}
