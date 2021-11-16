package models

type ViewApps struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"column:name"`
	UserId int    `json:"userId" gorm:"column:user_id"`
}

func (ViewApps) TableName() string {
	return "v_users_apps"
}
