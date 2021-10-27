package handlers

import (
	"Authentication-Service/gen/models"
	"Authentication-Service/gen/restapi/operations"
	"Authentication-Service/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

type LoginEmployee struct {}



func Lgoin() operations.LoginHandler{
	return &LoginEmployee{}
}


func (l LoginEmployee) Handle(params operations.LoginParams) middleware.Responder {
	token , error := service.Login(swag.StringValue(params.Login.Email) , swag.StringValue(params.Login.Password))
	if token == ""{
		return operations.NewLoginInternalServerError().WithPayload(error)
	}else {
		return operations.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
	}
}