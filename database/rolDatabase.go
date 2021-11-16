package database

import (
	"errors"
	"os"

	"github.com/jeisson-u/web-project-rol.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RolDatabase struct {
	db *gorm.DB
}

func (instance *RolDatabase) Start() (bool, error) {
	databaseActionConnection := os.Getenv("DATABASE_CONNECTION")
	db, err := gorm.Open(mysql.Open(databaseActionConnection), &gorm.Config{})
	if err != nil {
		return false, err
	}
	instance.db = db
	return true, nil
}

func (instance *RolDatabase) Close() {
	sqlDB, err := instance.db.DB()

	if err == nil {
		sqlDB.Close()
	}

}
func (instance *RolDatabase) GetAll() (*[]models.Rols, error) {

	var Rols []models.Rols

	result := instance.db.Where("del=?", 0).Where("state=?", 1).Find(&Rols)

	if result.Error != nil {
		return nil, result.Error
	}

	return &Rols, nil

}

func (instance *RolDatabase) Update(id int64, name string, password string, state int64) error {

	result := instance.db.Model(&models.Rols{}).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "state": state})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("register-dont-updated")
	}

	return nil

}

func (instance *RolDatabase) Delete(id int64) error {

	result := instance.db.Model(&models.Rols{}).Where("id = ?", id).Updates(map[string]interface{}{"del": 1})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("register-dont-deleted")
	}

	return nil

}

func (instance *RolDatabase) Create(name string, state int64) error {

	newRegister := models.Rols{Name: name, State: int(state), Del: false}
	result := instance.db.Create(&newRegister)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("rol-dont-create")
	}

	return nil

}
