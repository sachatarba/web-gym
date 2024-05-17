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
}

func NewClientMembershipRepo(db *gorm.DB) service.IClientMembershipsRepository {
	return &ClientMembershipRepo{
		db: db,
	}
}

func (r *ClientMembershipRepo) CreateNewClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	conv := orm.NewClientMembershipConverter()
	tx := r.db.Create(conv.ConvertFromEntity(clientMembership))

	return tx.Error
}

func (r *ClientMembershipRepo) ChangeClientMembership(ctx context.Context, clientMembership entity.ClientMembership) error {
	conv := orm.NewClientMembershipConverter()
	tx := r.db.WithContext(ctx).UpdateColumns(conv.ConvertFromEntity(clientMembership))

	return tx.Error
}
func (r *ClientMembershipRepo) DeleteClientMembership(ctx context.Context, clientMembershipID uuid.UUID) error {
	tx := r.db.WithContext(ctx).Delete(&orm.ClientMembership{
		ID: clientMembershipID,
	})

	return tx.Error
}
func (r *ClientMembershipRepo) GetClientMembershipByID(ctx context.Context, clientMembershipID uuid.UUID) (entity.ClientMembership, error) {
	membership := orm.ClientMembership{
		ID: clientMembershipID,
	}
	tx := r.db.WithContext(ctx).First(membership)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(service.ErrGetByIDNotFound, tx.Error)
	}

	if tx.Error != nil {
		return entity.ClientMembership{}, tx.Error
	}

	conv := orm.NewClientMembershipConverter()

	return conv.ConvertToEntity(membership), tx.Error
}

func (r *ClientMembershipRepo) ListClientMembershipsByClientID(ctx context.Context, clientID uuid.UUID) ([]entity.ClientMembership, error) {
	var memberships []orm.ClientMembership
	
	tx := r.db.WithContext(ctx).Where(&orm.Client{ID: clientID}).Find(&memberships)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx.Error = errors.Join(tx.Error, service.ErrListByIDNotFound)
	}

	if tx.Error != nil {
		return []entity.ClientMembership{}, tx.Error
	}

	conv := orm.NewClientMembershipConverter()
	clientMemberships := conv.ConvertToEntitySlice(memberships)

	return clientMemberships, tx.Error
}
