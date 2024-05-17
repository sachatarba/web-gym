package entity

type IConverter[Entity any, Converted any] interface {
	ConvertFromEntity(entity Entity) Converted
	ConvertToEntity(converted Converted) Entity

	ConvertFromEntitySlice(entity []Entity) []Converted
	ConvertToEntitySlice(converted []Converted) []Entity
}
