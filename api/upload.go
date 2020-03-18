package api

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type UploadRouter struct {
	*echo.Group
	Prefix string
	wf     WebFramework
}

func (r *UploadRouter) Test(c echo.Context) error {
	return c.String(http.StatusOK, "upload test")
}

func (r *UploadRouter) Upload(c echo.Context) error {
	label := c.FormValue("label")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		err := SaveToStorage(label, file)
		if err != nil {
			return c.String(http.StatusInternalServerError, "fail to save file")
		}

		// TODO: insert ppt db record
	}

	return c.String(http.StatusOK, "file upload success")
}

func (r *UploadRouter) Connect(w WebFramework) error {
	fmt.Println("connecting upload router")

	r.Group = w.Group(r.Prefix)
	r.wf = w
	r.GET("test", r.Test)
	r.POST("", r.Upload)
	return nil
}

const (
	StorageDirPath = "~/.marpurl/storage/"
	DefaultDirPerm = 0755
)

// TODO: abstract storage layer for using different storages
func SaveToStorage(label string, file *multipart.FileHeader) error {
	// SaveToStorage use localstorage temporary
	dirpath := StorageDirPath + label
	err := os.MkdirAll(dirpath, DefaultDirPerm)
	if os.IsExist(err) {
		fmt.Printf("%s is already exist\n", label)
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(dirpath + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func NewUploadRouter() *UploadRouter {
	r := &UploadRouter{
		Group:  nil,
		wf:     nil,
		Prefix: "/api/v1/upload",
	}
	return r
}

func Upload(c echo.Context) error {
	return c.String(http.StatusOK, "upload")
}
