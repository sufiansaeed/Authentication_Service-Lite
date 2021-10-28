package service

import (
	"Authentication-Service/db/MySqlClient"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var mySignKey = []byte("123123123123")

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(email string , password string) (string, int, string){
	databasePassword, userID, status:= MySqlClient.Login(email)
	if (status == "Error running database query") || (status == "Could not fetch password"){
		// 500
		return "", 500, "Internal Server Error"
	}else if status == "Could not find email"{
		//404
		return "", 404, "Couldn't find email"
	}else if status == "Password found" {
		loginSuccess := CheckPasswordHash(password, databasePassword)
		if loginSuccess == true {
			//200
			token, _ := GenerateJWT(email, userID)
			return token, 200, "Login successful"
		} else {
			status = "Incorrect password!"
			//401
			return "", 401, "Incorrect Password"
		}
	}
	//500
	return "", 500, "Internal Server Error"
}

func GenerateJWT(userEmail string, userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userEmail"] = userEmail
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()
	tokenString, err := token.SignedString(mySignKey)
	if err != nil {
		fmt.Println("Error is: ", err)
		return "", err
	}
	return tokenString, nil
}