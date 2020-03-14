package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UploadRouterMock struct {
	mock.Mock
}

func TestLoadUploadAPI(t *testing.T) {
	w := NewWebFramework()

	uploadRouter := NewUploadRouter("/api/upload")
	// 원하는 라우터만 연결할 수 있게 만들어야한다
	w.ConnectRouters(uploadRouter)
	// 라우터에 연결될 자원들을 인터페이스로 만들어서 목 테스트가 가능해야한다.
	// 여기서는 DB, Storage - 깨끗하게 초기화되어야한다. (테스트 환경으로 시작, 끝에)
	req := httptest.NewRequest(http.MethodGet, "/api/upload/test", nil)
	rec := httptest.NewRecorder()

	c := w.NewContext(req, rec)
	if assert.NoError(t, test(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "upload test", rec.Body.String())
	}
}
