package database

import (
	"os"

	"github.com/jeisson-u/web-project-rol.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppDatabase struct {
	db *gorm.DB
}

func (instance *AppDatabase) Start() (bool, error) {
	databaseActionConnection := os.Getenv("DATABASE_CONNECTION")
	db, err := gorm.Open(mysql.Open(databaseActionConnection), &gorm.Config{})
	if err != nil {
		return false, err
	}
	instance.db = db
	return true, nil
}

func (instance *AppDatabase) Close() {
	sqlDB, err := instance.db.DB()

	if err == nil {
		sqlDB.Close()
	}

}

func (instance *AppDatabase) GetByUser(userId int64) (*[]models.ViewApps, error) {

	var apps []models.ViewApps

	result := instance.db.Where("user_id=?", userId).Find(&apps)

	if result.Error != nil {
		return nil, result.Error
	}

	return &apps, nil

}
