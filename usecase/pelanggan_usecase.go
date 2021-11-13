package usecase

import (
	"bitbucket/models"
	"bitbucket/repository"
)

type IPelangganUsecase interface {
	SaveCustomer(cust *models.Pelanggan) (*models.Pelanggan, error)
}

type pelangganUsecaseRepo struct {
	custRepo 	repository.IPelangganRepository
}

func (p *pelangganUsecaseRepo) SaveCustomer(cust *models.Pelanggan) (*models.Pelanggan, error) {
	cust.Prepare()
	return p.custRepo.SaveCustomer(cust)
}

func NewCustomerUsecase(custRepo repository.IPelangganRepository ) IPelangganUsecase  {
	return &pelangganUsecaseRepo{custRepo,}
}