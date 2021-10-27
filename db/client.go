package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func SqlClient()  *sql.DB{
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "user"
	dbPort := 3306
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(127.0.0.1:%v)/%s", dbUser, dbPass, dbPort, dbName))
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
		fmt.Println("DB Error")
	}
	fmt.Println("DB Connection Successful!")
	return db
}

