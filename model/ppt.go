package model

import (
	"github.com/jinzhu/gorm"
)

type PPT struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(100);PRIMARY_KEY;NOT NULL"`
	DirPath string `json:"dirpath" gorm:"type:varchar(255);NOT NULL"`
}
