package handlers

import (
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
)

type SignupEmployee struct {}

func SignupEmployeeHandler() operations.EmployeeSignupHandler{
	return &SignupEmployee{}
}


func (s SignupEmployee) Handle(params operations.EmployeeSignupParams) middleware.Responder {
	_, err, code := service.SignupEmployee(GenModelToDomainModel(params.Body))

	if err != nil{
		fmt.Println("Error inserting into Database.")
		if code == 0 {
			return operations.NewEmployeeSignupInternalServerError()
		}else if code == 409{
			return operations.NewEmployeeSignupConflict()
		}
	}
	return operations.NewEmployeeSignupCreated()
}
