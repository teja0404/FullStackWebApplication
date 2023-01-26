package models

type Customer struct {
	Id        uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Gender    string `json:"gender"` //This hould be enum instead of string
	IsDeleted bool   `json:"isDeleted"`
	//EnrolledCourses []string `json:"enrolledCourses"`

	//Premium bool `json:"premium"`
}
