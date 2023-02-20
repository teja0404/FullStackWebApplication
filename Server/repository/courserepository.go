package repository

import (
	"net/http"
	"server/bo"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllCoursesFromDB(c *gin.Context) {

	var allCourses []models.Course
	initializers.DB.Find(&allCourses)
	c.JSON(http.StatusOK, gin.H{
		"allCourses": allCourses,
	})
}

func GetCourseByIdFromDB(id string, c *gin.Context) {
	var course models.Course
	initializers.DB.First(&course, "id = ?", id)
	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

func UpdateCourseInDBById(id string, c *gin.Context) {
	var course models.Course
	var courseBody bo.Course

	c.Bind(&courseBody)

	initializers.DB.First(&course, "id = ?", id)
	course.Name = courseBody.Name
	course.InstructorName = courseBody.InstructorName
	course.Price = courseBody.Price
	course.Description = courseBody.Description
	course.Duration = courseBody.Duration

	initializers.DB.Save(&course)
	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})

}

func DeleteCourseById(id string, c *gin.Context) {
	var course models.Course
	initializers.DB.Where("id = ?", id).Delete(&course)
}
