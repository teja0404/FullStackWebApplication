package models

type Admin struct {
	Id    uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
