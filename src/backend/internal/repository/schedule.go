package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
	"github.com/sachatarba/course-db/internal/orm"
	"github.com/sachatarba/course-db/internal/service"
	"gorm.io/gorm"
)

type ScheduleRepo struct {
	db        *gorm.DB
	converter entity.IConverter[entity.Schedule, orm.Schedule]
}

func NewScheduleRepo(db *gorm.DB) service.IScheduleRepository {
	return &ScheduleRepo{
		db:        db,
		converter: orm.NewScheduleConverter(),
	}
}

func (r *ScheduleRepo) CreateNewSchedule(ctx context.Context, schedule entity.Schedule) error {
	scheduleOrm := r.converter.ConvertFromEntity(schedule)
	tx := r.db.WithContext(ctx).Create(&scheduleOrm)

	return tx.Error
}

func (r *ScheduleRepo) ChangeSchedule(ctx context.Context, schedule entity.Schedule) error {
	scheduleOrm := r.converter.ConvertFromEntity(schedule)
	tx := r.db.WithContext(ctx).Save(&scheduleOrm)

	return tx.Error
}

func (r *ScheduleRepo) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Delete(&orm.Schedule{ID: scheduleID})

	return tx.Error
}

func (r *ScheduleRepo) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (entity.Schedule, error) {
	scheduleOrm := orm.Schedule{
		ID: scheduleID,
	}

	tx := r.db.WithContext(ctx).First(&scheduleOrm)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return entity.Schedule{}, service.ErrGetByIDNotFound
	}

	return r.converter.ConvertToEntity(scheduleOrm), tx.Error
}

func (r *ScheduleRepo) ListSchedulesByClientID(ctx context.Context, clientID uuid.UUID) ([]entity.Schedule, error) {
	var scheduleOrms []orm.Schedule
	tx := r.db.WithContext(ctx).
		Where(&orm.Schedule{ClientID: clientID}).
		Preload("Training").
		Find(&scheduleOrms)

	return r.converter.ConvertToEntitySlice(scheduleOrms), tx.Error
}
