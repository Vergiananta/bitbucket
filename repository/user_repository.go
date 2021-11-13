package repository

import (
	"bitbucket/middlewares"
	"bitbucket/models"
	"bitbucket/models/dto"
	"errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	SaveUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteAUser(uid string) error
	LoginUser(login *dto.LoginRequest) (string, error)
	FindUserByID(uid string) (*models.User, error)

}

type userRepositoryImpl struct {
	db *gorm.DB
}

func (u *userRepositoryImpl) LoginUser(login *dto.LoginRequest) (string, error) {
	var err error
	var accounts dto.LoginResponse

	if err = u.db.Debug().Raw("select * from users where email = ? ", login.User).Scan(&accounts).Error; err != nil {
		return "", err
	}
	err = models.VerifyPassword(accounts.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	if gorm.ErrRecordNotFound == err || accounts.Name == ""  {
		return "", errors.New("User Not Found")
	}
	return middlewares.CreateToken(&accounts)
}

func (u *userRepositoryImpl) SaveUser(newUser *models.User) (*models.User, error) {
	var err error
	var user models.User
	if err = u.db.Debug().Table("users").Where("email = ?", newUser.Email).First(&user).Error; user.Name != "" {
		return nil, errors.New("email has been registered")
	}
	if err = u.db.Debug().Table("users").Where("phone = ?", newUser.Phone).First(&user).Error; user.Name != "" {
		return nil, errors.New("phone number has been registered")
	}
	if err = u.db.Debug().Create(&newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userRepositoryImpl) UpdateUser(user *models.User) (*models.User, error) {
	u.db = u.db.Debug().Save(user)
	if err := u.db.Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepositoryImpl) DeleteAUser(id string)  error {
	uid, _ := uuid.FromString(id)
	var err error
	err= u.db.Debug().Raw("delete from users where id = ?", uid).Scan(&models.User{}).Error
	if err != nil {
		return err
	}
	if gorm.ErrRecordNotFound == err {
		return errors.New("User Not Found")
	}

	return nil
}

func (u *userRepositoryImpl) FindUserByID(id string) (*models.User, error) {
	var err error
	uid, _ := uuid.FromString(id)
	var user models.User
	err = u.db.Debug().Table("users").Where("id  = ?", uid).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.ErrRecordNotFound == err  || user.Name == ""{
		return nil, errors.New("User Not Found")
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) IUserRepository  {
	return &userRepositoryImpl{
		db,
	}
}