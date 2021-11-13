package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type TransaksiDetail struct {
	ID 				uuid.UUID	`gorm:"type:uuid;unique;index" json:"id"`
	Quantity 		int64		`gorm:"not null; column:quantity" json:"quantity"`
	ProductID		uuid.UUID	`gorm:"column:product_id" json:"product_id"`
	Product 		Product		`gorm:"foreignKey:ProductID;references:ID"`
	TransaksiID 	uuid.UUID 	`gorm:"column:transaksi_id" json:"transaksi_id"`
	Transaksi 		Transaksi	`gorm:"foreignKey:TransaksiID;references:ID"`
	CreatedAt     	time.Time   `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt     	time.Time   `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" `
	DeletedAt     	*time.Time  `gorm:"column:deleted_at" sql:"index"`
}

func (td *TransaksiDetail) Prepare() error  {
	td.ID = uuid.NewV4()
	td.CreatedAt = time.Now()
	return nil
}