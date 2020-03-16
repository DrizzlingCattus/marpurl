package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"marpurl/db"
	"marpurl/env"
	"marpurl/model"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UploadRouterMock struct {
	mock.Mock
}

type UploadDBTestSuite struct {
	suite.Suite
	DB       *gorm.DB
	Fixtures *testfixtures.Loader
}

func (suite *UploadDBTestSuite) SetupSuite() {
	env.Load(env.TestEnvFilename)
	suite.DB = db.Connect()
	suite.DB.AutoMigrate(&model.PPT{})

	sdb := suite.DB.DB()
	fixtures, err := testfixtures.New(
		testfixtures.Database(sdb),           // You database connection
		testfixtures.Dialect("mysql"),        // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("../fixture"), // the directory containing the YAML files
	)
	if err != nil {
		log.Fatalf("fixture setting fail %v", err)
	}
	suite.Fixtures = fixtures
	fmt.Println("setup suite")
}

func (suite *UploadDBTestSuite) TestUploadDB() {
	suite.Fixtures.Load()

	var ppt model.PPT
	suite.DB.Where("name = ?", "afile").First(&ppt)
	assert.Equal(suite.T(), ppt.DirPath, "/test/a")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UploadDBTestSuite))
}

func TestAcceptUploadAPI(t *testing.T) {
	uploadRouter := NewUploadRouter("/api/upload")
	// 라우터에 연결될 자원들을 인터페이스로 만들어서 Mock 테스트가 가능해야한다.

	// 원하는 라우터만 연결할 수 있게 만들어야한다
	w := NewWebFramework()
	w.ConnectRouters(uploadRouter)
	req := httptest.NewRequest(http.MethodGet, "/api/upload/test", nil)
	rec := httptest.NewRecorder()

	c := w.NewContext(req, rec)
	if assert.NoError(t, test(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "upload test", rec.Body.String())
	}

	// 작업이 끝마친 후 DB, Storage - 깨끗하게 초기화되어야한다. (테스트 환경으로 시작, 끝에)
}
