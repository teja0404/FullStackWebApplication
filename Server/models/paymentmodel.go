package models

type Payment struct {
	Id      uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name    string `json:"name"`
	Courses string `json:"courses"`
	Bill    int    `json:"bill"`
	Date    string
}
