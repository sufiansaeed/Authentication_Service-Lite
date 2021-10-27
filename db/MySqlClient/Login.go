package MySqlClient

import (
	client "Authentication-Service/db"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func Login(email string, password string) (bool, string, int){
	db := 	client.SqlClient()
	var databasePassword string
	var userID int
	getPasswordQuery := "SELECT password, ID FROM employeeSignup WHERE Email = '"+email+"'"
	result, err := db.Query(getPasswordQuery)
	if err != nil{
		return false, "Error with database connection", -1
	}
	for result.Next() {
		err := result.Scan(&databasePassword, &userID)
		if err != nil {
			fmt.Println("Error is: ", err)
		}
	}
	if databasePassword == ""{
		return false,"Email Not Found", -1
	} else {
		loginsuccess := (CheckPasswordHash(password, databasePassword))
		return loginsuccess, "", userID
	}
}