package orm

import "github.com/google/uuid"

type Equipment struct {
	ID          uuid.UUID `gorm:"type:UUID;primaryKey"`
	Name        string    `gorm:"type:TEXT"`
	Description string    `gorm:"type:TEXT"`
	GymID       uuid.UUID
}
