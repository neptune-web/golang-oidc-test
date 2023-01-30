package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Testdata struct
type Testdata struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	RequestID int       `json:"request_id" gorm:"unique"`
	// Requestername  string    `json:"requestername" gorm:"unique`
	// Expirationdate string    `json:"expirationdate"`
}

// Testdatas struct
type Testdatas struct {
	Testdatas []Testdata `json:"testdatas"`
}

func (testdata *Testdata) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	testdata.ID = uuid.New()
	return
}
