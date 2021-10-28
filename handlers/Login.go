package handlers

import (
	"Authentication-Service/gen/models"
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type LoginEmployee struct {}



func Login() operations.LoginHandler{
	return &LoginEmployee{}
}


func (l LoginEmployee) Handle(params operations.LoginParams) middleware.Responder {
	token , code, error := service.Login(swag.StringValue(params.Login.Email) , swag.StringValue(params.Login.Password))
	if token == ""{
		if code == 500 {
			return operations.NewLoginInternalServerError().WithPayload(error)
		}
		if code == 404{
			return operations.NewLoginNotFound().WithPayload(error)
		}
		if code == 401{
			return operations.NewLoginUnauthorized().WithPayload(error)
		}
	}else {
		return operations.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
	}
	return operations.NewLoginInternalServerError().WithPayload(error)
}