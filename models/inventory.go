package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Inventory struct {
	ID 			uuid.UUID 	`gorm:"unique; type:uuid; index" json:"id"`
	Quantity 	int64		`gorm:"not null" json:"quantity"`
	OutletID 	uuid.UUID	`gorm:"not null" json:"outlet_id"`
	Outlet 		Outlet		`gorm:"foreignKey:OutletID;references:ID"`
	ProductID 	uuid.UUID	`gorm:"not null" json:"product_id"`
	Product		Product		`gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 	*time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (i *Inventory) Prepare() error {
	i.ID = uuid.NewV4()
	i.CreatedAt = time.Now()
	return nil
}
