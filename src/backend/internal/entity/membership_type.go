package entity

import (
	"strconv"

	"github.com/google/uuid"
)

type MembershipType struct {
	ID           uuid.UUID
	Type         string
	Description  string
	Price        string
	DaysDuration int
	GymID        uuid.UUID
}

func (m *MembershipType) Validate() bool {
	if m.DaysDuration <= 0 {
		return false
	}

	if price, err := strconv.ParseFloat(m.Price, 64); price <= 0 || err != nil {
		return false
	}

	return m.Type != "" && m.Description != ""
}
