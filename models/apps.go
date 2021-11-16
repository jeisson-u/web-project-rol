package models

type Apps struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

func (Apps) TableName() string {
	return "apps"
}
