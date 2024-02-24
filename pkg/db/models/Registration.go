package models

import "time"

type Entity interface {
	GetId() int
}

const ModelRegistration = "registration"

type Registration struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Patron    string    `json:"patron"`
	Phone     Phone     `json:"phone"`
	Email     Email     `json:"email"`
}

type Email string
type Phone string

func (a *Registration) GetId() int {
	return a.Id
}
