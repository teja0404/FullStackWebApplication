package repository

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllCoursesFromDB(c *gin.Context) {

	var allCourses []models.Course

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Find(&allCourses)

	c.JSON(http.StatusOK, gin.H{
		"allCourses": allCourses,
	})
}

func GetCourseByIdFromDB(id string, c *gin.Context) {
	var course models.Course

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&course, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

func UpdateCourseInDBById(id string, c *gin.Context) {
	var course models.Course

	var Body struct {
		Name           string
		InstructorName string
		Price          int
		Description    string
		Duration       int
	}

	c.Bind(&Body)

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&course, "id = ?", id)
	course.Name = Body.Name
	course.InstructorName = Body.InstructorName
	course.Price = Body.Price
	course.Description = Body.Description
	course.Duration = Body.Duration

	initializers.DB.Save(&course)

	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})

}

func DeleteCourseById(id string, c *gin.Context) {
	var course models.Course

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Where("id = ?", id).Delete(&course)
}
