package api

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	Connector interface {
		Connect(w WebFramework) error
	}

	WebFramework interface {
		GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route

		NewContext(r *http.Request, w http.ResponseWriter) Context
		Use(middleware ...MiddlewareFunc)
		Pre(middleware ...MiddlewareFunc)
		Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
		Group(prefix string, m ...MiddlewareFunc) *echo.Group
	}

	webFramework struct {
		*echo.Echo
		DB *gorm.DB
	}

	HandlerFunc    = echo.HandlerFunc
	MiddlewareFunc = echo.MiddlewareFunc
	Route          = echo.Route
	Context        = echo.Context
	Group          = echo.Group
)

func NewWebFramework(db *gorm.DB) *webFramework {
	return &webFramework{
		Echo: echo.New(),
		// TODO: abstract db layer for loose coupling
		DB: db,
	}
}

func (w *webFramework) ConnectRouters(routers ...Connector) {
	for _, r := range routers {
		r.Connect(w)
	}
}
