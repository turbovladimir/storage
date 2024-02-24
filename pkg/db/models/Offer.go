package models

import "time"

const ModelOffer = "offer"

type Offer struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	AddedAt  time.Time `json:"added_at"`
	Name     string    `json:"name"`
	IsActive bool      `json:"is_active"`
	Img      string    `json:"img"`
}
