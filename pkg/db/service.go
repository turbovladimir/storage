package db

import (
	"fmt"
	"log"

	"github.com/turbovladimir/storage.git/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Database string `envconfig:"name"`
}

func NewConnect(config *Config) *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", config.User, config.Password, config.Database)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&models.Registration{}); err != nil {
		log.Fatalln(err)
	}

	return db
}
