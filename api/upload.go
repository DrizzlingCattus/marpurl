package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type UploadRouter struct {
	*echo.Group
	prefix string
	wf     WebFramework
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "upload test")
}

func (r *UploadRouter) Connect(w WebFramework) error {
	fmt.Println("connecting upload router")

	r.Group = w.Group(r.prefix)
	r.wf = w
	r.GET("test", test)
	return nil
}

func NewUploadRouter(prefix string) *UploadRouter {
	r := &UploadRouter{
		Group:  nil,
		wf:     nil,
		prefix: prefix,
	}
	return r
}

func Upload(c echo.Context) error {
	return c.String(http.StatusOK, "upload")
}
