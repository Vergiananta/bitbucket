package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Outlet struct {
	ID 			uuid.UUID	`gorm:"type:uuid;unique;index" json:"id"`
	Name 		string		`gorm:"not null; column:name" json:"name"`
	Address 	string		`gorm:"not null; column: address" json:"address"`
	MerchantID 	uuid.UUID	`gorm:"null" json:"merchant_id"`
	Merchant 	*Merchant 	`gorm:"foreignKey:MerchantID;references:ID;not null" json:"merchant"`
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 	*time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (o *Outlet) Prepare() error  {
	o.ID = uuid.NewV4()
	o.CreatedAt = time.Now()
	return nil
}