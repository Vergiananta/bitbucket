package models

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type Product struct {
	ID 			uuid.UUID	`gorm:"unique;index;type:uuid" json:"id" `
	Name 		string		`gorm:"not null" json:"name"`
	Deskripsi 	string 		`gorm:"not null" json:"deskripsi"`
	Price		int64		`gorm:"not null" json:"price"`
	CreatedAt    time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 	*time.Time 	`gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (u *Product) Prepare() error {
	u.ID = uuid.NewV4()
	u.CreatedAt = time.Now()
	return nil
}

func (u *Product) ValidateProduct(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Product Name")
		}
		if u.Deskripsi == "" {
			return errors.New("Required Description")
		}
		if u.Price == 0 {
			return errors.New("Required Price")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Product Name")
		}
		if u.Deskripsi == "" {
			return errors.New("Required Description")
		}
		if u.Price == 0 {
			return errors.New("Required Price")
		}
		return nil
	}

}
