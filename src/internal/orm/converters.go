package orm

import (
	"strconv"

	"github.com/sachatarba/course-db/internal/entity"
)

type MembershipTypeConverter struct {
}

func NewMembershipTypeConverter() entity.IConverter[entity.MembershipType, MembershipType] {
	return &MembershipTypeConverter{}
}

func (conv *MembershipTypeConverter) ConvertFromEntity(membership entity.MembershipType) MembershipType {
	return MembershipType{
		ID:           membership.ID,
		Type:         membership.Type,
		Description:  membership.Description,
		Price:        membership.Price,
		DaysDuration: membership.DaysDuration,
		GymID:        membership.GymID,
	}
}

func (conv *MembershipTypeConverter) ConvertToEntity(membership MembershipType) entity.MembershipType {
	return entity.MembershipType{
		ID:           membership.ID,
		Type:         membership.Type,
		Description:  membership.Description,
		Price:        membership.Price,
		DaysDuration: membership.DaysDuration,
		GymID:        membership.GymID,
	}
}

type ClientMembershipConverter struct {
}

func NewClientMembershipConverter() entity.IConverter[entity.ClientMembership, ClientMembership] {
	return &ClientMembershipConverter{}
}

func (conv *ClientMembershipConverter) ConvertFromEntity(membership entity.ClientMembership) ClientMembership {
	return ClientMembership{
		ID:               membership.ID,
		StartDate:        membership.StartDate,
		EndDate:          membership.EndDate,
		MembershipTypeID: membership.ID,
		MembershipType:   NewMembershipTypeConverter().ConvertFromEntity(membership.MembershipType),
		ClientID:         membership.ClientID,
	}
}

func (c *ClientMembershipConverter) ConvertToEntity(membership ClientMembership) entity.ClientMembership {
	return entity.ClientMembership{
		ID:             membership.ID,
		StartDate:      membership.StartDate,
		EndDate:        membership.EndDate,
		MembershipType: NewMembershipTypeConverter().ConvertToEntity(membership.MembershipType),
		ClientID:       membership.ClientID,
	}
}

type EquipmentConverter struct {
}

func NewEquipmentConverter() entity.IConverter[entity.Equipment, Equipment] {
	return &EquipmentConverter{}
}

func (conv *EquipmentConverter) ConvertFromEntity(e entity.Equipment) Equipment {
	return Equipment{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		GymID:       e.GymID,
	}
}

func (conv *EquipmentConverter) ConvertFromEntitySlice(e []entity.Equipment) []Equipment {
	eqps := make([]Equipment, 0, len(gym.Equipments))
			converter := NewEquipmentConverter()
			for i := 0; i < len(gym.Equipments); i++ {
				eqps[i] = converter.ConvertFromEntity(gym.Equipments[i])
			}

	return eqps
}

func (conv *EquipmentConverter) ConvertToEntity(e Equipment) entity.Equipment {
	return entity.Equipment{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		GymID:       e.GymID,
	}
}

type GymConverter struct {
}

func NewGymConverter() entity.IConverter[entity.Gym, Gym] {
	return &GymConverter{}
}

func (conv *GymConverter) ConvertFromEntity(gym entity.Gym) Gym {
	return Gym{
		ID: gym.ID,
		Name: gym.Name,
		Phone: gym.Phone,
		City: gym.City,
		Addres: gym.Addres,
		IsChain: strconv.FormatBool(gym.IsChain),
		Trainers: gym.Trainers,
		Equipments: func() []Equipment{
			eqps := make([]Equipment, 0, len(gym.Equipments))
			converter := NewEquipmentConverter()
			for i := 0; i < len(gym.Equipments); i++ {
				eqps[i] = converter.ConvertFromEntity(gym.Equipments[i])
			}

			return eqps
		}(),
		: func() []Equipment{
			eqps := make([]Equipment, 0, len(gym.Equipments))
			converter := NewEquipmentConverter()
			for i := 0; i < len(gym.Equipments); i++ {
				eqps[i] = converter.ConvertFromEntity(gym.Equipments[i])
			}

			return eqps
		}(),
	}
}

func (conv *GymConverter) ConvertToEntity(gym Gym) entity.Gym {
	return entity.Equipment{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		GymID:       e.GymID,
	}
}
