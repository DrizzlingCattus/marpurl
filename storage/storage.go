package storage

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type (
	Metadata map[string]string

	Storage interface {
		Save(key string, meta *Metadata, reader io.Reader) error
		Load(key string) (*Metadata, *bytes.Buffer)
	}

	FileStorage struct {
		Rootpath       string
		SavedFileNames []string
	}
)

//func (fst *FileStorage) Save(key string, meta *Metadata, reader io.Reader) error {
//	_, err := os.Stat(fst.path(key))
//	if os.IsNotExist(err) {
//	} else if os.IsExist(err) {
//
//	}
//	return err
//}
//
//func (fst *FileStorage) saveFile() error {
//}
//
//func (fst *FileStorage) path(filename string) string {
//	return fst.Rootpath + filename
//}
//
//func (fst *FileStorage) Load(key string) (*Metadata, *bytes.Buffer) {
//
//}

func makeStorageDir(rootname string) (string, error) {
	home, _ := os.UserHomeDir()
	rootpath := filepath.Join(home, rootname)

	// TODO: check err except isExist, isNotExist
	os.MkdirAll(rootpath, 0755)
	return rootpath, nil
}

func NewFileStorage() *FileStorage {
	rootpath, _ := makeStorageDir(".marpurl")

	infos, err := ioutil.ReadDir(rootpath)
	if err != nil {
		log.Fatal("FileStorage cannot read rootpath infos", err)
	}

	sfns := make([]string, 50, 200)
	for i, info := range infos {
		sfns[i] = info.Name()
	}

	return &FileStorage{
		Rootpath:       rootpath,
		SavedFileNames: sfns,
	}
}
