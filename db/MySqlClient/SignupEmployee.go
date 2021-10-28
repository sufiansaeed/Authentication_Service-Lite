package MySqlClient

import (
	client "Authentication-Service/db"
	"Authentication-Service/models"
	"fmt"
)



func SignupEmployee(employee *models.DomainEmployee) (int, string){
	db := 	client.SqlClient()
	checkQuery := "SELECT * FROM employeeSignup WHERE Email = '"+employee.Email+"'"
	insertQuery := "INSERT INTO employeeSignup (FullName, FatherName, Age, Email, Password, Gender, NIC) VALUES ('" + employee.FullName + "', '" + employee.FatherName + "', '" + fmt.Sprintf("%v",employee.Age) +"', '" + employee.Email + "', '" + employee.Password + "', '" + employee.Gender + "', '"+employee.NIC+"')"

	result, err := db.Query(checkQuery)
	if err != nil{
		return 500, "Error With Database Connection"
	}
	if result.Next(){
		return 409, "Email Already Exists"

	}else{
		result, err := db.Query(insertQuery)
		if err != nil{
			return 500, "Error Inserting Values in Database"
		}
		defer result.Close()
		return 201, "Employee Created Successfully!"
	}
}

