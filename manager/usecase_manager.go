package manager

import (
	"bitbucket/connect"
	"bitbucket/usecase"
)

type ServiceManager interface {
	UserUseCase() usecase.IUserUseCase
	ProductUseCase() usecase.IProductUsecase
}

type serviceManager struct {
	repo RepoManager
}

func (sm *serviceManager) ProductUseCase() usecase.IProductUsecase {
	return usecase.NewProductUsecase(sm.repo.ProductRepo())
}

func (sm *serviceManager) UserUseCase() usecase.IUserUseCase {
	return usecase.NewUserUseCase(sm.repo.UserRepo())
}

func NewServiceManager(connect connect.Connect) ServiceManager{
	return &serviceManager{repo: NewRepoManager(connect)}
}

