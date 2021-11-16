package models

type UsersRols struct {
	Id     int `json:"id" gorm:"primaryKey"`
	UserId int `json:"userId" gorm:"column:user_id"`
	RolId  int `json:"rolId" gorm:"column:rol_id"`
}

func (UsersRols) TableName() string {
	return "users_rols"
}
