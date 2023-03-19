package models

import (
	"gorm.io/gorm"
)

type Live struct {
	gorm.Model
	Name         string `gorm:"name"`
	Owner        string `gorm:"owner"`
	RecordStatus string `gorm:"record_status"`
	RecordPath   string `gorm:"record_path"`
}
