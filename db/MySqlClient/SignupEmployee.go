package MySqlClient

import (
	client "Authentication-Service/db"
	"Authentication-Service/models"
	"fmt"
)

func SignupEmployee(employee *models.DomainEmployee) (string, error, int){
	db := 	client.SqlClient()
	checkQuery := "SELECT * FROM employeeSignup WHERE Email = '"+employee.Email+"'"
	insertQuery := "INSERT INTO employeeSignup (FullName, FatherName, Age, Email, Password, Gender, NIC) VALUES ('" + employee.FullName + "', '" + employee.FatherName + "', '" + fmt.Sprintf("%v",employee.Age) +"', '" + employee.Email + "', '" + employee.Password + "', '" + employee.Gender + "', '"+employee.NIC+"')"

	result, err := db.Query(checkQuery)
	if err != nil{
		fmt.Println(" mySqlClient: Error with database connection")
		fmt.Println(err)
		return "", err, 0
	}
	if result.Next(){
		fmt.Println("c")
		fmt.Println(err)
		err := fmt.Errorf("Entered Email already exists!")
		return "", err, 409

	}else{
		result, err := db.Query(insertQuery)
		if err != nil{
			fmt.Println(" mySqlClient: Error inserting values in database")
			return "Error!", err, 0
		}
		fmt.Println("Data inserted into database!")
		defer result.Close()
		return "", nil, 0
	}
}

