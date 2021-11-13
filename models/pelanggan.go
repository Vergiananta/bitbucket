package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Pelanggan struct {
	ID 				uuid.UUID	`json:"id"`
	Name 			string		`json:"name"`
	Address 		string		`json:"address"`
	PhoneNumber 	int64		`json:"phone_number"`
	CreatedAt    time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (p *Pelanggan) Prepare() error  {
	p.ID = uuid.NewV4()
	p.CreatedAt = time.Now()
	return nil
}