package config

import (
	"backend/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := Appconfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to init database %v",err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(Appconfig.Database.maxIdleConns)
	sqlDB.SetMaxOpenConns(Appconfig.Database.maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("failed to config database %v",err)
	}

	global.DB = db
}
