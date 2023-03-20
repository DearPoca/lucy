package models

import (
	"gorm.io/gorm"
)

type Live struct {
	gorm.Model
	Name         string `gorm:"name"`
	Owner        string `gorm:"owner"`
	WebrtcUrl    string `gorm:"webrtc_url"`
	RtmpUrl      string `gorm:"rtmp_url"`
	HttpFlvUrl   string `gorm:"http_flv_url"`
	RecordStatus string `gorm:"record_status"`
	RecordPath   string `gorm:"record_path"`
}
