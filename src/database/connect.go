package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"os"

	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connection to database")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("failed conection to database")
	}
	db.SetConnMaxIdleTime(100)
	db.SetMaxOpenConns(10)

	return gormDB, nil
}
