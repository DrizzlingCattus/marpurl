package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func ConnectDB() *gorm.DB {
	DBMS := "mysql"
	DBUser := "user"
	DBPassword := "62445"
	DBName := "marpurl"

	DBConnectOptions := fmt.Sprintf("%s:%s@(localhost:3306)/", DBUser, DBPassword)
	db, err := gorm.Open(DBMS, DBConnectOptions)
	if err != nil {
		log.Fatal("Fail to connect DB", err)
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

type ContextWithDB struct {
	echo.Context
	DB *gorm.DB
}

type Cat struct {
	gorm.Model
	Name string
	Type string
}

type PPT struct {
	gorm.Model
	Name    string
	DirPath string
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func main() {
	db := ConnectDB()
	defer db.Close()

	db.AutoMigrate(&PPT{})

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cdb := &ContextWithDB{Context: c, DB: db}
			return next(cdb)
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
