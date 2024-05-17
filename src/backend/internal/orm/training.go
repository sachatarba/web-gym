package orm

import "github.com/google/uuid"

type Training struct {
	ID           uuid.UUID `gorm:"type:UUID;primaryKey"`
	Title        string    `gorm:"type:TEXT"`
	Description  string    `gorm:"type:TEXT"`
	TrainingType string    `sql:"type:ENUM('aerobic', 'anaerobic', 'flexibility', 'strength')"`
	TrainerID    uuid.UUID
}
