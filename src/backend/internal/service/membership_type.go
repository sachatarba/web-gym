package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type MembershipTypeService struct {
	membershipTypeRepo IMembershipTypeRepository
}

func NewMembershipTypeService(repository IMembershipTypeRepository) IMembershipTypeService {
	return &MembershipTypeService{
		membershipTypeRepo: repository,
	}
}

func (s *MembershipTypeService) RegisterNewMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	if !membershipType.Validate() {
		return ErrValidation
	}

	err := s.membershipTypeRepo.RegisterNewMembershipType(ctx, membershipType)

	return err
}

func (s *MembershipTypeService) ChangeMembershipType(ctx context.Context, membershipType entity.MembershipType) error {
	if !membershipType.Validate() {
		return ErrValidation
	}

	err := s.membershipTypeRepo.ChangeMembershipType(ctx, membershipType)
	return err
}

func (s *MembershipTypeService) DeleteMembershipType(ctx context.Context, membershipTypeID uuid.UUID) error {
	err := s.membershipTypeRepo.DeleteMembershipType(ctx, membershipTypeID)

	return err
}

func (s *MembershipTypeService) GetMembershipTypeByID(ctx context.Context, membershipTypeID uuid.UUID) (entity.MembershipType, error) {
	membershipType, err := s.membershipTypeRepo.GetMembershipTypeByID(ctx, membershipTypeID)
	if err != nil {
		return entity.MembershipType{}, err
	}

	return membershipType, nil
}

func (s *MembershipTypeService) ListMembershipTypesByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.MembershipType, error) {
	membershipTypes, err := s.membershipTypeRepo.ListMembershipTypesByGymID(ctx, gymID)
	if err != nil {
		return []entity.MembershipType{}, err
	}

	return membershipTypes, nil
}
