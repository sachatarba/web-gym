package entity

import (
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/pkg/validator"
)

type Trainer struct {
	ID            uuid.UUID
	Fullname      string
	Email         string
	Phone         string
	Qualification string
	UnitPrice     float64
	GymsID        []uuid.UUID
	Trainings     []Training
}

func (t *Trainer) Validate() bool {
	if !validator.IsValidEmail(t.Email) {
		return false
	}

	if !validator.IsValidPhoneNumber(t.Phone) {
		return false
	}

	if t.UnitPrice <= 0 {
		return false
	}

	return t.Fullname != "" && t.Qualification != ""
}
