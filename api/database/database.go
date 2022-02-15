package database

import (
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetUpDB() (conn *gorm.DB, err error) {
	conn, err = gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db, err := conn.DB()
	if err != nil {
		db.Close()
	}

	return conn, err
}