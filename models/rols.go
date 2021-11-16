package models

type Rols struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"column:name"`
	State int    `json:"state" gorm:"column:state"`
	Del   bool   `json:"del" gorm:"column:del"`
}

func (Rols) TableName() string {
	return "rols"
}
