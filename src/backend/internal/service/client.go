package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type ClientService struct {
	clientRepo IClientRepository
}

func NewClientService(clientRepo IClientRepository) IClientService {
	return &ClientService{
		clientRepo: clientRepo,
	}
}

func (s *ClientService) RegisterNewClient(ctx context.Context, client entity.Client) error {
	if !client.Validate() {
		return ErrValidation
	}

	err := s.clientRepo.RegisterNewClient(ctx, client)

	return err
}

func (s *ClientService) ChangeClient(ctx context.Context, client entity.Client) error {
	if !client.Validate() {
		return ErrValidation
	}

	err := s.clientRepo.ChangeClient(ctx, client)

	return err
}

func (s *ClientService) DeleteClient(ctx context.Context, clientID uuid.UUID) error {
	err := s.clientRepo.DeleteClient(ctx, clientID)

	return err
}

func (s *ClientService) GetClientByID(ctx context.Context, clientID uuid.UUID) (entity.Client, error) {
	client, err := s.clientRepo.GetClientByID(ctx, clientID)
	if err != nil {
		return entity.Client{}, err
	}

	return client, nil
}

func (s *ClientService) GetClientByLogin(ctx context.Context, login string) (entity.Client, error) {
	client, err := s.clientRepo.GetClientByLogin(ctx, login)
	if err != nil {
		return entity.Client{}, err
	}

	return client, nil
}

func (s *ClientService) ListClients(ctx context.Context) ([]entity.Client, error) {
	clients, err := s.clientRepo.ListClients(ctx)
	if err != nil {
		return []entity.Client{}, err
	}

	return clients, err
}
