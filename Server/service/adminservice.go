package service

import (
	"net/http"
	"server/repository"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "Fullstack Application BE is running",
	})

}

type Requestbody struct {
	Name           string
	InstructorName string
	Price          int
	Description    string
	Duration       int
}

func AddCourse(c *gin.Context) {
	var requestbody Requestbody
	c.Bind(&requestbody)
	repository.AddCourseInDB(requestbody.Name, requestbody.InstructorName, requestbody.Price, requestbody.Description, requestbody.Duration, c)
}
