package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Store struct
type Store struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;"`
	StoreID int       `json:"storeid" gorm:"unique"`
}

// Stores struct
type Stores struct {
	Stores []Store `json:"stores"`
}

func (store *Store) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	store.ID = uuid.New()
	return
}
