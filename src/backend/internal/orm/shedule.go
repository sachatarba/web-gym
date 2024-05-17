package orm

import "github.com/google/uuid"

type Schedule struct {
	ID           uuid.UUID `gorm:"type:UUID;primaryKey"`
	DayOfTheWeek string    `gorm:"type:DATE"`
	StartTime    string    `gorm:"type:TIME"`
	EndTime      string    `gorm:"type:TIME"`
	ClientID     uuid.UUID
	TrainingID   uuid.UUID
	Training     Training
}
