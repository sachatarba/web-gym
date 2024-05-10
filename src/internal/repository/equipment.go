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

type EquipmentRepo struct {
	db *gorm.DB
	// CreateNewEquipment(ctx context.Context, equipment entity.Equipment) error
	// ChangeEquipment(ctx context.Context, equipment entity.Equipment) error
	// DeleteEquipment(ctx context.Context, equipmentID uuid.UUID) error
	// GetEquipmentByID(ctx context.Context, equipmentID uuid.UUID) (entity.Equipment, error)
	// ListEquipmentsByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.Equipment, error)
}

func NewEquipmentRepo(db *gorm.DB) service.IEquipmentRepository {
	return &EquipmentRepo{
		db: db,
	}
}

func (r *EquipmentRepo) CreateNewEquipment(ctx context.Context, equipment entity.Equipment) error {
	equip := orm.Equipment{
		ID:          equipment.ID,
		Name:        equipment.Name,
		Description: equipment.Description,
		GymID:       equipment.GymID,
	}
	tx := r.db.Create(equip)

	return tx.Error
}

func (r *EquipmentRepo) ChangeEquipment(ctx context.Context, equipment entity.Equipment) error {
	equip := orm.Equipment{
		ID:          equipment.ID,
		Name:        equipment.Name,
		Description: equipment.Description,
		GymID:       equipment.GymID,
	}
	tx := r.db.UpdateColumns(equip)

	return tx.Error
}
func (r *EquipmentRepo) DeleteEquipment(ctx context.Context, equipmentID uuid.UUID) error {
	tx := r.db.Delete(&orm.Equipment{
		ID: equipmentID,
	})

	return tx.Error
}
func (r *EquipmentRepo) GetEquipmentByID(ctx context.Context, equipmentID uuid.UUID) (entity.Equipment, error) {
	equipment := &orm.Equipment{
		ID: equipmentID,
	}
	tx := r.db.First(equipment)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(service.ErrGetByIDNoSuchRow, tx.Error)
	}

	if tx.Error != nil {
		return entity.Equipment{}, tx.Error
	}

	return entity.Equipment{
		ID:          equipment.ID,
		Name:        equipment.Name,
		Description: equipment.Description,
		GymID:       equipment.GymID,
	}, tx.Error
}

func (r *EquipmentRepo) ListEquipmentsByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.Equipment, error) {
	gym := &orm.Gym{
		ID: gymID,
	}
	tx := r.db.Preload("Equipments").First(gym)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(tx.Error, service.ErrListByIDNoRows)
	}

	if tx.Error != nil {
		return []entity.Equipment{}, tx.Error
	}

	equipments := make([]entity.Equipment, len(gym.Equipments))
	for i := 0; i < len(gym.Equipments); i++ {
		equipments[i] = entity.Equipment{
			ID:          gym.Equipments[i].ID,
			Name:        gym.Equipments[i].Name,
			Description: gym.Equipments[i].Description,
			GymID:       gym.Equipments[i].GymID,
		}
	}

	return equipments, tx.Error
}
