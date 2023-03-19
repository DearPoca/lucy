package models

import (
	"fmt"

	"lucy/pkg/log"

	"lucy/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	var err error
	loggers := logger.New(log.Logger, logger.Config{})
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.User,
		setting.MysqlSetting.Password,
		setting.MysqlSetting.Host,
		setting.MysqlSetting.Database)

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{Logger: loggers})

	if err != nil {
		log.Fatal("gorm.Open failed", "err", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Database connect failed", "err", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(-1)

	db.AutoMigrate(&User{})
}

func Db() *gorm.DB {
	return db
}
