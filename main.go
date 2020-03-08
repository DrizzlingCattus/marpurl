package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var (
	DBMS       string
	DBUser     string
	DBPassword string
	DBName     string
)

func InitEnv(mode string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBMS = os.Getenv("DBMS")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
}

func ConnectDB() *gorm.DB {
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

type PPT struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(100);PRIMARY_KEY;NOT NULL"`
	DirPath string `json:"dirpath" gorm:"type:varchar(255);NOT NULL"`
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func main() {
	// TODO: Mode = os.arg
	var Mode string = "dev"
	InitEnv(Mode)

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

	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":8080"))
}
