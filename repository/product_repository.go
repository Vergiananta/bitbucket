package repository

import (
	"bitbucket/models"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IProductRepository interface {
	SaveProduct(newProduct *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	FindByIdProduct(id string) (*models.Product, error)
	DeleteProduct(id string) error
	GetProductByPaginate(page, size string) ([]*models.Product, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func (p *productRepositoryImpl) FindByIdProduct(id string) (*models.Product, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var product models.Product
	err = p.db.Debug().Table("products").Where("id  = ?", uid).Find(&product).Error
	if err != nil {
		return nil, err
	}
	if gorm.ErrRecordNotFound == err {
		return nil, errors.New("product Not Found")
	}
	return &product, nil
}

func (p *productRepositoryImpl) DeleteProduct(id string)  error {
	uid, _ := uuid.FromString(id)
	var err error
	err= p.db.Debug().Raw("delete from products where id = ?", uid).Scan(&models.Product{}).Error
	if err != nil {
		return err
	}
	if gorm.ErrRecordNotFound == err  {
		return errors.New("product Not Found")
	}

	return nil
}

func (p *productRepositoryImpl) SaveProduct(newProduct *models.Product) (*models.Product, error) {
	var err error
	if err = p.db.Debug().Create(&newProduct).Error; err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (p *productRepositoryImpl) UpdateProduct(product *models.Product) (*models.Product, error) {
	p.db = p.db.Debug().Save(product)
	if err := p.db.Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepositoryImpl) GetProductByPaginate(page, size string) ([]*models.Product, error) {
	offset, pageSize := Paginate(page, size)
	var err error
	products := make([]*models.Product,0)
	if err = p.db.Debug().Raw("select * from products offset ? limit ?",  offset, pageSize).
		Scan(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}


func NewProductRepository(db *gorm.DB) IProductRepository  {
	return &productRepositoryImpl{
		db,
	}
}
