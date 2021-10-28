package MySqlClient

import (
	client "Authentication-Service/db"
)



func Login(email string) (string, int, string) {
	db := client.SqlClient()
	var databasePassword string
	var userID int
	getPasswordQuery := "SELECT password, ID FROM employeeSignup WHERE Email = '" + email + "'"
	result, err := db.Query(getPasswordQuery)
	if err != nil {
		return "", -1, "Error running database query"
	}
	for result.Next() {
		err := result.Scan(&databasePassword, &userID)
		if err != nil {
			return "", -1, "Could not fetch password"
		}
		return databasePassword, userID, "Password found"
	}
	return "", -1, "Could not find email"
}