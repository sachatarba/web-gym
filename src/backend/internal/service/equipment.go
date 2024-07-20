package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type EquipmentService struct {
	equipmentRepo IEquipmentRepository
}

func NewEquipmentService(equipmentRepo IEquipmentRepository) IEquipmentService {
	return &EquipmentService{
		equipmentRepo: equipmentRepo,
	}
}

func (s *EquipmentService) CreateNewEquipment(ctx context.Context, equipment entity.Equipment) error {
	if !equipment.Validate() {
		return ErrValidation
	}

	err := s.equipmentRepo.CreateNewEquipment(ctx, equipment)

	return err
}

func (s *EquipmentService) ChangeEquipment(ctx context.Context, equipment entity.Equipment) error {
	if !equipment.Validate() {
		return ErrValidation
	}

	err := s.equipmentRepo.ChangeEquipment(ctx, equipment)

	return err
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, equipmentID uuid.UUID) error {
	err := s.equipmentRepo.DeleteEquipment(ctx, equipmentID)

	return err
}

func (s *EquipmentService) GetEquipmentByID(ctx context.Context, equipmentID uuid.UUID) (entity.Equipment, error) {
	equipment, err := s.equipmentRepo.GetEquipmentByID(ctx, equipmentID)
	if err != nil {
		return entity.Equipment{}, err
	}

	return equipment, nil
}

func (s *EquipmentService) ListEquipmentsByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.Equipment, error) {
	equipments, err := s.equipmentRepo.ListEquipmentsByGymID(ctx, gymID)
	if err != nil {
		return []entity.Equipment{}, err
	}

	return equipments, nil
}
