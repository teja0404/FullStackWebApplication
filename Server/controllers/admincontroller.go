package controllers

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "Fullstack Application BE is running",
	})

}

func AddCourse(c *gin.Context) {

	var requestbody struct {
		Name           string
		InstructorName string
		Price          int
		Description    string
		Duration       int
	}

	c.Bind(&requestbody)

	course := models.Course{Name: requestbody.Name, InstructorName: requestbody.InstructorName, Price: requestbody.Price, Description: requestbody.Description, Duration: requestbody.Duration}

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	result := initializers.DB.Create(&course)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}
