package database

import (
	"errors"
	"os"

	"github.com/jeisson-u/web-project-rol.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserDatabase struct {
	db *gorm.DB
}

func (instance *UserDatabase) Start() (bool, error) {
	databaseActionConnection := os.Getenv("DATABASE_CONNECTION")
	db, err := gorm.Open(mysql.Open(databaseActionConnection), &gorm.Config{})
	if err != nil {
		return false, err
	}
	instance.db = db
	return true, nil
}

func (instance *UserDatabase) Close() {
	sqlDB, err := instance.db.DB()

	if err == nil {
		sqlDB.Close()
	}

}
func (instance *UserDatabase) GetAll() (*[]models.Users, error) {

	var users []models.Users

	result := instance.db.Where("del=?", 0).Where("state=?", 1).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil

}
func (instance *UserDatabase) GetByEmailAndPassword(email string, password string) (*models.Users, error) {

	var user models.Users

	result := instance.db.Where("email=?", email).Where("password=?", password).Where("del=?", 0).Where("state=?", 1).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}
func (instance *UserDatabase) Search(content string) (*[]models.Users, error) {

	var users []models.Users

	result := instance.db.Where("del=?", 0).Where("state=?", 1).Where("email like ?", "%"+content+"%").Or("name like ?", "%"+content+"%").First(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil

}
func (instance *UserDatabase) Update(id int64, name string, password string, state int64) error {

	result := instance.db.Model(&models.Users{}).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "password": password, "state": state})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("register-dont-updated")
	}

	return nil

}

func (instance *UserDatabase) Delete(id int64) error {

	result := instance.db.Model(&models.Users{}).Where("id = ?", id).Updates(map[string]interface{}{"del": 1})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("register-dont-deleted")
	}

	return nil

}

func (instance *UserDatabase) Create(name string, email string, password string, state int64) error {

	newRegister := models.Users{Email: email, Password: password, Name: name, State: int(state), Del: false}
	result := instance.db.Create(&newRegister)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user-dont-create")
	}

	return nil

}
