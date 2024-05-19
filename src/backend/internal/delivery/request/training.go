package request

import "github.com/google/uuid"

type Training struct {
	ID           uuid.UUID
	Title        string
	Description  string
	TrainingType string
	TrainerID    uuid.UUID
}
