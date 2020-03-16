package main

import (
	"net/http"

	"marpurl/db"
	"marpurl/env"
	"marpurl/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type ContextWithDB struct {
	echo.Context
	DB *gorm.DB
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func main() {
	// TODO: Mode = os.arg
	env.Load(env.DevEnvFilename)

	tdb := db.Connect()
	defer tdb.Close()

	tdb.AutoMigrate(&model.PPT{})

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
