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

type ClientMembershipRepo struct {
	db *gorm.DB
	// CreateNewClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error
	// ChangeClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error
	// DeleteClientMembership(ctx context.Context, clientMembershipID uuid.UUID) error
	// GetClientMembershipByID(ctx context.Context, clientMembershipID uuid.UUID) (entity.ClientMembership, error)
	// ListClientMembershipsByClientID(ctx context.Context, clientID uuid.UUID) ([]entity.ClientMembership, error)
}

func NewClientMembershipRepo(db *gorm.DB) service.IClientMembershipsRepository {
	return &ClientMembershipRepo{
		db: db,
	}
}

func (r *ClientMembershipRepo) CreateNewClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	membership := orm.ClientMembership{
		ID:               clientMembership.ID,
		StartDate:        clientMembership.StartDate,
		EndDate:          clientMembership.EndDate,
		MembershipTypeID: clientMembership.MembershipType.ID,
		ClientID:         clientMembership.ClientID,
	}
	tx := r.db.Create(membership)

	return tx.Error
}

func (r *ClientMembershipRepo) ChangeClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	membership := orm.ClientMembership{
		ID:               clientMembership.ID,
		StartDate:        clientMembership.StartDate,
		EndDate:          clientMembership.EndDate,
		MembershipTypeID: clientMembership.MembershipType.ID,
		ClientID:         clientMembership.ClientID,
	}
	tx := r.db.UpdateColumns(membership)

	return tx.Error
}
func (r *ClientMembershipRepo) DeleteClientMembership(ctx context.Context, clientMembershipID uuid.UUID) error {
	tx := r.db.Delete(&orm.ClientMembership{
		ID: clientMembershipID,
	})

	return tx.Error
}
func (r *ClientMembershipRepo) GetClientMembershipByID(ctx context.Context, clientMembershipID uuid.UUID) (entity.ClientMembership, error) {
	membership := &orm.ClientMembership{
		ID: clientMembershipID,
	}
	tx := r.db.First(membership)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(service.ErrGetByIDNoSuchRow, tx.Error)
	}

	if tx.Error != nil {
		return entity.ClientMembership{}, tx.Error
	}

	return entity.ClientMembership{
		ID:             membership.ID,
		StartDate:      membership.StartDate,
		EndDate:        membership.EndDate,
		MembershipType: entity.MembershipType{
			ID:
		},
	}, tx.Error
}

func (r *ClientMembershipRepo) ListClientMembershipsByGymID(ctx context.Context, gymID uuid.UUID) ([]entity.ClientMembership, error) {
	gym := &orm.Gym{
		ID: gymID,
	}
	tx := r.db.Preload("ClientMemberships").First(gym)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(tx.Error, service.ErrListByIDNoRows)
	}

	if tx.Error != nil {
		return []entity.ClientMembership{}, tx.Error
	}

	ClientMemberships := make([]entity.ClientMembership, len(gym.ClientMemberships))
	for i := 0; i < len(gym.ClientMemberships); i++ {
		ClientMemberships[i] = entity.ClientMembership{
			ID:          gym.ClientMemberships[i].ID,
			Name:        gym.ClientMemberships[i].Name,
			Description: gym.ClientMemberships[i].Description,
			GymID:       gym.ClientMemberships[i].GymID,
		}
	}

	return ClientMemberships, tx.Error
}
