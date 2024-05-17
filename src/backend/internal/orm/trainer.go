package orm

import "github.com/google/uuid"

type Trainer struct {
	ID            uuid.UUID `gorm:"type:UUID;primaryKey"`
	Fullname      string    `gorm:"type:TEXT"`
	Email         string    `gorm:"type:TEXT;check:check_valid_email, (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}$')"`
	Phone         string    `gorm:"type:TEXT;check:check_valid_international_phone, (phone ~ '^\\+[0-9]+-[0-9]+-[0-9]+-[0-9]+-[0-9]+')"`
	Qualification string    `gorm:"type:TEXT"`
	UnitPrice     float64   `gorm:"type:REAL"`
	Gyms          []*Gym    `gorm:"many2many:gym_trainers;"`
	Trainings     []Training
}
