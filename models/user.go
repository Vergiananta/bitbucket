package models

import (
	"errors"
	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID 			uuid.UUID 	`gorm:"type:uuid; unique; index; column:id;" json:"id"`
	Name 		string 		`gorm:"not null" json:"name"`
	Email 		string		`gorm:"not null; unique; column:email" json:"email"`
	Address 	string 		`gorm:"not null; column:address" json:"address"`
	Phone 		int64 		`gorm:"not null; unique; column:phone" json:"phone"`
	Password 	string 		`gorm:"not null; column:password" json:"password"`
	Merchant 	*[]Merchant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"merchant"`
	CreatedAt    time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt 	*time.Time 	`gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (u *User) Prepare() error {
	u.ID = uuid.NewV4()
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()
	return nil
}

func (u *User) EditUser() error  {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.UpdatedAt = time.Now()
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		if u.Address == "" {
			return errors.New("Required Address")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		if u.Address == "" {
			return errors.New("Required Address")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}

}