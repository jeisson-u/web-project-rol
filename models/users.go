package models

import "time"

type Users struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	Name      string     `json:"name" gorm:"column:name"`
	State     int        `json:"state" gorm:"column:state"`
	Del       bool       `json:"del" gorm:"column:del"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Users) TableName() string {
	return "users"
}
