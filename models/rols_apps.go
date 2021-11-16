package models

type RolsApps struct {
	Id    int `json:"id" gorm:"primaryKey"`
	RolId int `json:"rolId" gorm:"column:rol_id"`
	AppId int `json:"appId" gorm:"column:app_id"`
}

func (RolsApps) TableName() string {
	return "rols_apps"
}
