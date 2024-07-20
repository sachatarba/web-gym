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

type GymRepo struct {
	db        *gorm.DB
	converter entity.IConverter[entity.Gym, orm.Gym]
}

func NewGymRepo(db *gorm.DB) service.IGymRepository {
	return &GymRepo{
		db:        db,
		converter: orm.NewGymConverter(),
	}
}

func (r *GymRepo) RegisterNewGym(ctx context.Context, gym entity.Gym) error {
	gymOrm := r.converter.ConvertFromEntity(gym)
	tx := r.db.Create(&gymOrm)

	return tx.Error
}

func (r *GymRepo) ChangeGym(ctx context.Context, gym entity.Gym) error {
	gymOrm := r.converter.ConvertFromEntity(gym)
	tx := r.db.WithContext(ctx).Save(&gymOrm)

	return tx.Error
}

func (r *GymRepo) DeleteGym(ctx context.Context, gymID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Delete(&orm.Gym{
		ID: gymID},
	)

	return tx.Error
}

func (r *GymRepo) GetGymByID(ctx context.Context, gymID uuid.UUID) (entity.Gym, error) {
	gymOrm := orm.Gym{
		ID: gymID,
	}

	tx := r.db.WithContext(ctx).First(&gymOrm)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return entity.Gym{}, service.ErrGetByIDNotFound
	}

	return r.converter.ConvertToEntity(gymOrm), tx.Error
}

func (r *GymRepo) ListGyms(ctx context.Context) ([]entity.Gym, error) {
	var gymOrms []orm.Gym
	tx := r.db.WithContext(ctx).Find(&gymOrms)

	return r.converter.ConvertToEntitySlice(gymOrms), tx.Error
}
