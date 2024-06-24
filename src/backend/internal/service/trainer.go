package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type TrainerService struct {
	trainerRepo ITrainerRepository
}

func NewTrainerService(repository ITrainerRepository) ITrainerService {
	return &TrainerService{
		trainerRepo: repository,
	}
}

func (s *TrainerService) RegisterNewTrainer(ctx context.Context, trainer entity.Trainer) error {
	if !trainer.Validate() {
		return ErrValidation
	}

	err := s.trainerRepo.RegisterNewTrainer(ctx, trainer)

	return err
}

func (s *TrainerService) ChangeTrainer(ctx context.Context, trainer entity.Trainer) error {
	if !trainer.Validate() {
		return ErrValidation
	}

	err := s.trainerRepo.ChangeTrainer(ctx, trainer)

	return err
}

func (s *TrainerService) DeleteTrainer(ctx context.Context, trainerID uuid.UUID) error {
	err := s.trainerRepo.DeleteTrainer(ctx, trainerID)

	return err
}

func (s *TrainerService) GetTrainerByID(ctx context.Context, trainerID uuid.UUID) (entity.Trainer, error) {
	trainer, err := s.trainerRepo.GetTrainerByID(ctx, trainerID)
	if err != nil {
		return entity.Trainer{}, err
	}

	return trainer, nil
}

func (s *TrainerService) ListTrainers(ctx context.Context) ([]entity.Trainer, error) {
	trainers, err := s.trainerRepo.ListTrainers(ctx)
	if err != nil {
		return []entity.Trainer{}, err
	}

	return trainers, nil
}

func (s *TrainerService) ListTrainersByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.Trainer, error) {
	trainers, err := s.trainerRepo.ListTrainersByGymID(ctx, gymID)
	if err != nil {
		return []entity.Trainer{}, err
	}

	return trainers, nil
}
