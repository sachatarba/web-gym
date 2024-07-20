package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
	"github.com/sachatarba/course-db/internal/orm"
	"github.com/sachatarba/course-db/internal/service"
	"gorm.io/gorm"
)

type TrainingRepo struct {
	db        *gorm.DB
	converter entity.IConverter[entity.Training, orm.Training]
}

func NewTrainingRepo(db *gorm.DB) service.ITrainingRepository {
	return &TrainingRepo{
		db:        db,
		converter: orm.NewTrainingConverter(),
	}
}

func (r *TrainingRepo) CreateNewTraining(ctx context.Context, training entity.Training) error {
	trainingOrm := r.converter.ConvertFromEntity(training)
	tx := r.db.WithContext(ctx).Create(&trainingOrm)

	return tx.Error
}

func (r *TrainingRepo) ChangeTraining(ctx context.Context, training entity.Training) error {
	trainingOrm := r.converter.ConvertFromEntity(training)
	tx := r.db.WithContext(ctx).Save(&trainingOrm)

	return tx.Error
}

func (r *TrainingRepo) DeleteTraining(ctx context.Context, trainingID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Delete(&orm.Training{ID: trainingID})

	return tx.Error
}

func (r *TrainingRepo) ListTrainingsByTrainerID(ctx context.Context, trainerID uuid.UUID) ([]entity.Training, error) {
	var trainingOrms []orm.Training
	tx := r.db.WithContext(ctx).Where(&orm.Training{TrainerID: trainerID}).Find(&trainingOrms)

	return r.converter.ConvertToEntitySlice(trainingOrms), tx.Error
}
