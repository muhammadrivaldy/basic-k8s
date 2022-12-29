package models

type User struct {
	Id      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

func (u User) TableName() string {
	return "mst_user"
}
