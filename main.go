package main

import (
	"fmt"
	"log"
	"net/http"

	"marpurl/db"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

const (
	ModeDev = iota
	ModeProd
	ModeTest
)

func InitEnv(mode int) {
	if mode == ModeDev {
		fmt.Println("Loading Development envs")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
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
	InitEnv(ModeDev)

	tdb := db.Connect()
	defer tdb.Close()

	tdb.AutoMigrate(&PPT{})

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cdb := &ContextWithDB{Context: c, DB: tdb}
			return next(cdb)
		}
	})

	e.GET("/test", test)

	e.Logger.Fatal(e.Start(":8080"))
}
