package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "postgresql://neondb_owner:npg_iX3EVpTKBjP0@ep-steep-dream-a7ptpf74-pooler.ap-southeast-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  DSN,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			PrepareStmt: false,
		},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
