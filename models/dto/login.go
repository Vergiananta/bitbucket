package dto

import uuid "github.com/satori/go.uuid"

type LoginRequest struct {
	User 		string 	`json:"user"`
	Password 	string	`json:"password"`
}

type LoginResponse struct {
	ID 			uuid.UUID 	`json:"id"`
	Name 		string 		`json:"name"`
	Email 		string		`json:"email"`
	Address 	string 		`json:"address"`
	Phone 		string 		`json:"phone"`
	Password 	string		`json:"password"`
}