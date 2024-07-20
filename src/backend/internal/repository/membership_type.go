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

type MembershipTypeRepo struct {
	db        *gorm.DB
	converter entity.IConverter[entity.MembershipType, orm.MembershipType]
}

func NewMembershipTypeRepo(db *gorm.DB) service.IMembershipTypeRepository {
	return &MembershipTypeRepo{
		db:        db,
		converter: orm.NewMembershipTypeConverter(),
	}
}

func (r *MembershipTypeRepo) RegisterNewMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	membershipTypeOrm := r.converter.ConvertFromEntity(membershipType)
	tx := r.db.WithContext(ctx).Create(&membershipTypeOrm)

	return tx.Error
}

func (r *MembershipTypeRepo) ChangeMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	membershipTypeOrm := r.converter.ConvertFromEntity(membershipType)
	tx := r.db.WithContext(ctx).Save(&membershipTypeOrm)

	return tx.Error
}

func (r *MembershipTypeRepo) DeleteMembershipType(ctx context.Context, membershipTypeID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Delete(&orm.MembershipType{ID: membershipTypeID})

	return tx.Error
}

func (r *MembershipTypeRepo) GetMembershipTypeByID(ctx context.Context, membershipTypeID uuid.UUID) (entity.MembershipType, error) {
	membershipTypeOrm := orm.MembershipType{
		ID: membershipTypeID,
	}

	tx := r.db.WithContext(ctx).First(&membershipTypeOrm)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return entity.MembershipType{}, service.ErrGetByIDNotFound
	}

	return r.converter.ConvertToEntity(membershipTypeOrm), tx.Error
}

func (r *MembershipTypeRepo) ListMembershipTypesByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.MembershipType, error) {
	gym := &orm.Gym{
		ID: gymID,
	}
	tx := r.db.WithContext(ctx).Preload("MembershipTypes").First(gym)
	membershipTypeOrms := gym.MembershipTypes

	return r.converter.ConvertToEntitySlice(membershipTypeOrms), tx.Error
}
