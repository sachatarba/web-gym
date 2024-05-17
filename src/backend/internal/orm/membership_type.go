package orm

import "github.com/google/uuid"

type MembershipType struct {
	ID           uuid.UUID `gorm:"type:UUID;primaryKey"`
	Type         string    `gorm:"type:TEXT"`
	Description  string    `gorm:"type:TEXT"`
	Price        string    `gorm:"type:REAL"`
	DaysDuration int       `gorm:"type:INT"`
	GymID        uuid.UUID
}
