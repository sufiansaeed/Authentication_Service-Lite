package service

import (
	"Authentication-Service/db/MySqlClient"
	"Authentication-Service/models"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func SignupEmployee(employee *models.DomainEmployee) (int, string){
	employee.Password, _ = HashPassword(employee.Password)
	return MySqlClient.SignupEmployee(employee)
}