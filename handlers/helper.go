package handlers

import (
	model "Authentication-Service/gen/models"
	Domain "Authentication-Service/models"
	"github.com/go-openapi/swag"
)

func GenModelToDomainModel(employee *model.Employee) *Domain.DomainEmployee {
	return &Domain.DomainEmployee{
		FullName: swag.StringValue(employee.FullName),
		FatherName: swag.StringValue(employee.FatherName),
		Age: int(swag.Int64Value(employee.Age)),
		Email: swag.StringValue(employee.Email),
		Password: swag.StringValue(employee.Password),
		NIC: swag.StringValue(employee.NIC),
		Gender: swag.StringValue(employee.Gender),
	}
}


func DomainModelToGenModel(employee *Domain.DomainEmployee) * model.Employee{
	return &model.Employee{
		FullName: swag.String(employee.FullName),
		FatherName: swag.String(employee.FatherName),
		Age: swag.Int64(int64(employee.Age)),
		Email: swag.String(employee.Email),
		Password: swag.String(employee.Password),
		NIC: swag.String(employee.NIC),
		Gender: swag.String(employee.Gender),
	}
}