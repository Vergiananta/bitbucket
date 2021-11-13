package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Merchant struct {
	ID 		uuid.UUID	`gorm:"type:uuid;unique;index" json:"id"`
	Npwp 	string		`gorm:"column:npwp" json:"npwp"`
	Address string 		`gorm:"column:bussiness_address" json:"address"`
	Phone 	int64  		`gorm:"column:bussiness_phone" json:"phone"`
	Outlet 	[]Outlet	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"outlet"`
	UserID 	uuid.UUID 	`gorm:"null"`
	User 	User 		`gorm:"foreignKey:UserID; references:ID; not null" json:"user"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (m *Merchant) Prepare()  {
	m.ID = uuid.NewV4()
	m.CreatedAt = time.Now()
}