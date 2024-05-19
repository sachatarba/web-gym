package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type ClientMembershipsService struct {
	clientMembershipRepo IClientMembershipsRepository
}

func NewClientMembershipService(clientMembershipRepo IClientMembershipsRepository) IClientMembershipsService {
	return &ClientMembershipsService{
		clientMembershipRepo: clientMembershipRepo,
	}
}

func (m *ClientMembershipsService) CreateNewClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	if !clientMembership.Validate() {
		return ErrValidation
	}

	err := m.clientMembershipRepo.CreateNewClientMembership(ctx, clientMembership)
	return err
}

func (s *ClientMembershipsService) ChangeClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	if !clientMembership.Validate() {
		return ErrValidation
	}

	err := s.clientMembershipRepo.ChangeClientMembership(ctx, clientMembership)

	return err
}

func (s *ClientMembershipsService) DeleteClientMembership(ctx context.Context, clientMembershipID uuid.UUID) error {
	err := s.clientMembershipRepo.DeleteClientMembership(ctx, clientMembershipID)

	return err
}

func (s *ClientMembershipsService) GetClientMembershipByID(ctx context.Context, clientMembershipID uuid.UUID) (entity.ClientMembership, error) {
	membership, err := s.clientMembershipRepo.GetClientMembershipByID(ctx, clientMembershipID)
	if err != nil {
		return entity.ClientMembership{}, err
	}

	return membership, nil
}

func (s *ClientMembershipsService) ListClientMembershipsByClientID(ctx context.Context, clientID uuid.UUID) ([]entity.ClientMembership, error) {
	memberships, err := s.clientMembershipRepo.ListClientMembershipsByClientID(ctx, clientID)
	if err != nil {
		return []entity.ClientMembership{}, err
	}

	return memberships, nil
}
