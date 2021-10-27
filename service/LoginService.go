package service

import (
	"Authentication-Service/db/MySqlClient"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var mySignKey = []byte("123123123123")

func Login(email string , password string) (string,string){

	loginSuccess, errorString, userID := MySqlClient.Login(email, password)
	if loginSuccess == false && errorString == ""{
		errorString = "You have entered the wrong password."
	}
	if loginSuccess == true{
		token, _ := GenerateJWT(email, userID)
		return token ,"Login Successful and token generated successfully."
	}else {
		return "", errorString
	}
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