package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBMS       string
	DBUser     string
	DBPassword string
	DBName     string
)

func initEnv() {
	DBMS = os.Getenv("DBMS")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}

func Connect() *gorm.DB {
	initEnv()
	DBConnectOptions := fmt.Sprintf("%s:%s@(localhost:3306)/", DBUser, DBPassword)
	db, err := gorm.Open(DBMS, DBConnectOptions)
	if err != nil {
		log.Fatal("Fail to connect DB : ", err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS " + DBName)

	// option reference - https://github.com/go-sql-driver/mysql#parameters
	DBConnectOptions = fmt.Sprintf("%s:%s@(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBName)
	db, err = gorm.Open("mysql", DBConnectOptions)
	if err != nil {
		log.Fatal("Fail to connect DB", err)
	}
	return db
}
