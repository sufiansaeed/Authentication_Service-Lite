package handlers

import (
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"github.com/go-openapi/runtime/middleware"
)

type SignupEmployee struct {}

func SignupEmployeeHandler() operations.EmployeeSignupHandler{
	return &SignupEmployee{}
}


func (s SignupEmployee) Handle(params operations.EmployeeSignupParams) middleware.Responder {
	code, response := service.SignupEmployee(GenModelToDomainModel(params.Body))
	if code == 500{
		return operations.NewEmployeeSignupInternalServerError().WithPayload(response)
	}else if code == 409{
		return operations.NewEmployeeSignupConflict().WithPayload(response)
	}else if code == 201{
		return operations.NewEmployeeSignupCreated()
	}else{
		return operations.NewEmployeeSignupInternalServerError().WithPayload(response)
	}
}
