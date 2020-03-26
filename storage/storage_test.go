package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FileStorageSuite struct {
	suite.Suite
	fst *FileStorage
}

func (suite *FileStorageSuite) SetupSuite() {
	//t := suite.T()
	suite.fst = NewFileStorage()

}

func (suite *FileStorageSuite) TestHello() {

}

type OsSuite struct {
	suite.Suite
	currentWd string
	testDir   string
}

func (suite *OsSuite) SetupSuite() {
	suite.currentWd, _ = os.Getwd()
	suite.testDir = "test-marpurl"
}

func (suite *OsSuite) AfterTest(_, _ string) {
	err := os.Chdir(suite.currentWd)
	assert.NoError(suite.T(), err)
}

func (suite *OsSuite) TestMakeStorageDir() {
	t := suite.T()
	tmpdir := suite.testDir
	path, err := makeStorageDir(tmpdir)
	if assert.NoError(t, err) {
		home, _ := os.UserHomeDir()
		assert.Equal(t, path, filepath.Join(home, tmpdir))

		err = os.Chdir(path)
		assert.NoError(t, err)

		os.Remove(path)
	}
}

func (suite *OsSuite) TestMkdir() {
	t := suite.T()
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, suite.testDir)

	// there is no path file
	_, err := os.Stat(path)
	assert.Error(t, err)

	os.Mkdir(path, 0755)

	// there must exist path file
	_, err = os.Stat(path)
	assert.NoError(t, err)
}

func TestSuites(t *testing.T) {
	suite.Run(t, new(OsSuite))
	// suite.Run(t, new(FileStorageSuite))
}
