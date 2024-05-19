package postrgres_adapter

import (
	"gorm.io/gorm"
)

type IPostgresMigrator interface {
	Migrate() error
}

type PostgresMigrator struct {
	DB     *gorm.DB
	Tables []any
}

func (migrator *PostgresMigrator) Migrate() error {
	err := migrator.DB.AutoMigrate(
		migrator.Tables...,
	)

	return err
}
