package usecase

import (
	"bitbucket/models"
	"bitbucket/models/dto"
	"bitbucket/repository"
)

type IUserUseCase interface {
	Register(newUser *models.User) (*models.User, error)
	UpdateUser(editUser *models.User) (*models.User, error)
	LoginUser(login *dto.LoginRequest) (string, error)
	FindUserById(id string) (*models.User, error)
	DeleteUser(id string) error
}

type UserUsecaseRepo struct {
	userRepo repository.IUserRepository
}

func (u *UserUsecaseRepo) DeleteUser(id string) error {
	return u.userRepo.DeleteAUser(id)
}

func (u *UserUsecaseRepo) LoginUser(login *dto.LoginRequest) (string, error) {
	return u.userRepo.LoginUser(login)
}

func (u *UserUsecaseRepo) UpdateUser(editUser *models.User) (*models.User, error) {
	editUser.Validate("update")
	editUser.EditUser()
	return u.userRepo.UpdateUser(editUser)
}

func (u *UserUsecaseRepo) Register(newUser *models.User) (*models.User, error) {
	newUser.Validate("")
	newUser.Prepare()
	return u.userRepo.SaveUser(newUser)
}

func (u *UserUsecaseRepo) FindUserById(id string) (*models.User, error) {
	return u.userRepo.FindUserByID(id)
}

func NewUserUseCase(userRepo repository.IUserRepository) IUserUseCase {
	return &UserUsecaseRepo{
		userRepo,
	}
}
