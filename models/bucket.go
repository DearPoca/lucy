package models

import (
	"fmt"

	"lucy/pkg/setting"

	"gorm.io/gorm"
)

type Bucket struct {
	Name          string
	RootDirectory string
	Size          int64
	Capacity      int64
}

func CreateBucket(name string, cap int64) (bool, error) {
	var bucket Bucket
	err := db.Where(Bucket{Name: name}).First(&bucket).Error

	if err != gorm.ErrRecordNotFound {
		return false, fmt.Errorf("bucket已存在")
	}

	db.Create(&Bucket{
		Name:          name,
		RootDirectory: fmt.Sprintf("%s/%s", setting.AppSetting.AppRoot, name),
		Size:          0,
		Capacity:      cap,
	})

	return true, nil
}

func GetBucket(name string) (*Bucket, error) {
	var bucket Bucket
	err := db.Where(Bucket{Name: name}).First(&bucket).Error

	if err != nil {
		return nil, err
	}

	return &bucket, nil
}
