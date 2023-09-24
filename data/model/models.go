package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string

	Payments []Payment

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Payment struct {
	gorm.Model

	Name    string
	EventId string
	Price   int
	Payees  []User `gorm:"many2many:payees;"`
	Payer   User
	PayerId uint
}

type User struct {
	gorm.Model
	Name    string
	EventId string
	Event   Event
	Payments []Payment `gorm:"many2many:payees;"`
}
