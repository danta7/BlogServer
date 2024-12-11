package models

import "fmt"

// ImageModel 图片表
type ImageModel struct {
	Model
	Filename string `gorm:"size:64" json:"filename"`
	Path     string `gorm:"size:256" json:"path"`
	Size     int64  `json:"size"`
	Hash     string `gorm:"size:32" json:"hash"`
}

func (m ImageModel) WebPath() string {
	return fmt.Sprintf("/")
}