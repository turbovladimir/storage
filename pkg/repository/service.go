package repository

import (
	"fmt"
	"github.com/turbovladimir/storage.git/pkg/db"
	"github.com/turbovladimir/storage.git/pkg/models"
	"os"

	"gorm.io/gorm"
)

const (
	ErrorRecordNotFound = "record not found"
)

type Repository struct {
	DB     *gorm.DB
	entity interface{}
}

func CreateRegistrationRepo() *Repository {
	return new(db.NewConnect(&db.Config{os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")}), models.Registration{})
}

func CreateOfferRepo() *Repository {
	return new(db.NewConnect(&db.Config{os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")}), models.Offer{})
}

func new(db *gorm.DB, entity interface{}) *Repository {
	db = db.Model(entity)
	r := Repository{db, entity}

	return &r
}

func (r *Repository) Create(object interface{}) error {
	if result := r.DB.Create(object); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) GetBy(filter map[string][]string, target interface{}) {

	for field, val := range filter {
		val := val[0]
		r.DB.Where(fmt.Sprintf("%s = ?", field), val)
	}

	r.DB.Find(target)
}
