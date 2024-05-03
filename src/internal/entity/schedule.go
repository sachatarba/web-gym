package entity

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID            uuid.UUID
	DayOfTheeWeek string
	StartTime     string
	EndTime       string
	ClientID      uuid.UUID
	TrainingID    uuid.UUID
	Training      Training
}

func (sh *Schedule) Validate() bool {
	date, err := time.Parse(time.DateOnly, sh.DayOfTheeWeek)
	if err != nil {
		return false
	}
	sh.DayOfTheeWeek = date.String()

	startTime, err := time.Parse(time.DateTime, sh.StartTime)
	if err != nil {
		return false
	}
	sh.StartTime = startTime.String()

	endTime, err := time.Parse(time.DateTime, sh.EndTime)
	if err != nil {
		return false
	}
	sh.EndTime = endTime.String()

	return startTime.Before(endTime)
}
