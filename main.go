package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func TestGORM(db *gorm.DB) {
	// CRUD - create
	db.Create(&Cat{Name: "cat1", Type: "water"})
	db.Create(&Cat{Name: "cat2", Type: "water"})

	// CRUD - read
	var cat1 Cat
	db.First(&cat1, "type = ?", "water") // find cat , id 1
	err := fmt.Sprintf("cat1 %s %s\n", cat1.Name, cat1.Type)
	io.WriteString(os.Stdout, err)

	// CRUD - update
	db.Model(&cat1).Update("type", "fire")
	// check update
	var cat2 Cat
	db.First(&cat2, "type = ?", "water") // select from Cat where type = 'fire'
	err = fmt.Sprintf("cat2 %s %s\n", cat2.Name, cat2.Type)
	io.WriteString(os.Stdout, err)

	// CRUD - delete
	db.Delete(&cat1)
	db.Delete(&cat2)
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func main() {

	db := ConnectDB()
	defer db.Close()

	db.AutoMigrate(&Cat{})

	TestGORM(db)

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cdb := &ContextWithDB{Context: c, DB: db}
			return next(cdb)
		}
	})

	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":8080"))
}
