package dao

import (
	"blog/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DbConnection *gorm.DB

func InitDb(config configs.Config) {
	db, err := gorm.Open(postgres.Open(config.Postgres.PgConnString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect postgres: %s", err)
	}
	DbConnection = db
}
