package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InitializeDatabase struct {
	db *gorm.DB
}

func (instance *InitializeDatabase) Start() (bool, error) {
	databaseActionConnection := os.Getenv("DATABASE_CONNECTION")
	db, err := gorm.Open(mysql.Open(databaseActionConnection), &gorm.Config{})
	if err != nil {
		return false, err
	}
	instance.db = db
	return true, nil
}
func (instance *InitializeDatabase) Close() {
	sqlDB, err := instance.db.DB()

	if err == nil {
		sqlDB.Close()
	}

}

func (instance *InitializeDatabase) Create() error {

	// read file script
	pwd, _ := os.Getwd()
	dat, err := os.ReadFile(pwd + "/database/init.sql")

	if err != nil {
		return err
	}
	fmt.Println(string(dat))

	result := instance.db.Exec(string(dat))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
