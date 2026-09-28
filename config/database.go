package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(databaseURL string) (*gorm.DB, error) {
	return gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  databaseURL,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			PrepareStmt: false,
		},
	)
}
