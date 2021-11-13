package manager

import (
	"bitbucket/connect"
	"bitbucket/repository"
)

type RepoManager interface {
	UserRepo() repository.IUserRepository
	ProductRepo() repository.IProductRepository
}

type repoManager struct {
	connect connect.Connect
}

func (rm *repoManager) ProductRepo() repository.IProductRepository {
	return repository.NewProductRepository(rm.connect.SqlDb())
}

func (rm *repoManager) UserRepo() repository.IUserRepository {
	return repository.NewUserRepository(rm.connect.SqlDb())
}


func NewRepoManager(connect connect.Connect) RepoManager {
	return &repoManager{connect}
}