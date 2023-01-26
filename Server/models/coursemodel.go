package models

type Course struct {
	Id             uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name           string `json:"name"`
	InstructorName string `json:"instructorName"`
	Price          int    `json:"price"`
	Description    string `json:"description"`
	Duration       int    `json:"duration"`
	//EnrolledCount  int    `json:"enrolledCount"`
}
