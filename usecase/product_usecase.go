package usecase

import (
	"bitbucket/models"
	"bitbucket/repository"
	"time"
)

type IProductUsecase interface {
	SaveProduct(newProduct *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	FindByIdProduct(id string) (*models.Product, error)
	DeleteProduct(id string) error
	FindProductWithPaginate(page, size string) ([]*models.Product, error)
}

type productUsecase struct {
	productRepo repository.IProductRepository
}

func (p *productUsecase) FindProductWithPaginate(page, size string) ([]*models.Product, error) {
	return p.productRepo.GetProductByPaginate(page,size)
}

func (p *productUsecase) SaveProduct(newProduct *models.Product) (*models.Product, error) {
	newProduct.ValidateProduct("")
	newProduct.Prepare()
	return p.productRepo.SaveProduct(newProduct)
}

func (p *productUsecase) UpdateProduct(product *models.Product) (*models.Product, error) {
	product.ValidateProduct("update")
	product.UpdatedAt = time.Now()
	return p.productRepo.UpdateProduct(product)
}

func (p *productUsecase) FindByIdProduct(id string) (*models.Product, error) {
	return p.productRepo.FindByIdProduct(id)
}

func (p *productUsecase) DeleteProduct(id string)  error {
	return p.productRepo.DeleteProduct(id)
}

func NewProductUsecase(repo repository.IProductRepository) IProductUsecase  {
	return &productUsecase{
		repo,
	}
}