package repository

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddCourseInDB(Name string, InstructorName string, Price int, Description string, Duration int, c *gin.Context) {

	course := models.Course{Name: Name, InstructorName: InstructorName, Price: Price, Description: Description, Duration: Duration}

	initializers.DB = initializers.EstablishConnection()

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
