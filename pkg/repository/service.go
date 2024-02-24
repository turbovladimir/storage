package repository

import (
	"fmt"
	"github.com/turbovladimir/storage.git/pkg/db"
	"github.com/turbovladimir/storage.git/pkg/db/models"
	"os"
	"strings"

	"gorm.io/gorm"
)

type ModelRepository interface {
	Find(filter map[string]interface{}, models interface{})
	Create(params *CreateParams)
	Update(params *UpdateParams)
}

type Repository struct {
	DB *gorm.DB
}

var repo *Repository

func New() ModelRepository {
	if repo == nil {
		c := db.NewConnect(&db.Config{os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")})
		repo = &Repository{c}
	}

	return repo
}

func (r *Repository) Find(filter map[string]interface{}, m interface{}) {
	var tx *gorm.DB

	for field, expr := range filter {
		if strings.Contains(field, "_not") {
			field = strings.ReplaceAll(field, "_not", "")
			if field == "id" {
				tx = r.DB.Not(expr)
			}

		}
	}

	tx.Find(m)

	if err := tx.Error; err != nil {
		panic("Error occurrence when after scan repository.")
	}
}

func (r *Repository) Update(params *UpdateParams) {

}

func (r *Repository) Create(params *CreateParams) {
	result := r.DB.Model(models.Registration{}).Save(params.Models)

	params.Error = result.Error
}

func (r *Repository) GetBy(filter map[string][]string, target interface{}) {

	for field, val := range filter {
		val := val[0]
		r.DB.Where(fmt.Sprintf("%s = ?", field), val)
	}

	r.DB.Find(target)
}
