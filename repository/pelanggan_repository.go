package repository

import (
	"bitbucket/models"
	"gorm.io/gorm"
)

type IPelangganRepository interface {
	SaveCustomer(cust *models.Pelanggan) (*models.Pelanggan, error)
}

type custRepositoryImpl struct {
	db *gorm.DB
}

func (c *custRepositoryImpl) SaveCustomer(newCust *models.Pelanggan) (*models.Pelanggan, error) {
	var err error
	if err = c.db.Debug().Create(&newCust).Error; err != nil {
		return nil, err
	}
	return newCust, nil
}

func NewCustomerRepository(db *gorm.DB) IPelangganRepository {
	return &custRepositoryImpl{
		db,
	}
}
