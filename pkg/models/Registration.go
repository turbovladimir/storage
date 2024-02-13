package models

import "time"

type Entity interface {
	GetId() int
}

type Registration struct {
	Id      int       `json:"id" gorm:"primaryKey"`
	AddedAt time.Time `json:"added_at"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Patron  string    `json:"patron"`
	Phone   Phone     `json:"phone"`
	Email   Email     `json:"email"`
}

type Email string
type Phone string

func (a *Registration) GetId() int {
	return a.Id
}
