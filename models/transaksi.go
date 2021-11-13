package models

import uuid "github.com/satori/go.uuid"

type Transaksi struct {
	ID 					uuid.UUID			`gorm:"type:uuid; unique; index" json:"id"`
	TanggalTransaksi	string 				`gorm:"type:date; not null; column:tanggal_transaksi" json:"tanggal_transaksi"`
	TotalTransaksi 		int64 				`gorm:"not null; column:total_transaksi" json:"total_transaksi"`
	OutletID 			uuid.UUID			`gorm:"not null; column:outlet_id" json:"outlet_id"`
	Outlet				Outlet				`gorm:"foreignKey:OutletID;references:ID" `
	TransaksiDetail 	[]TransaksiDetail 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaksi_detail"`
}