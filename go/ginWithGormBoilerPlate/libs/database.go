package libs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	DatabaseHost := os.Getenv("DatabaseHost")
	DatabaseUser := os.Getenv("DatabaseUser")
	DatabasePassword := os.Getenv("DatabasePassword")
	DatabasePort := os.Getenv("DatabasePort")
	DatabaseName := os.Getenv("DatabaseName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseUser, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		db = nil
		panic("Failed to open database")
	}
	return db
}
