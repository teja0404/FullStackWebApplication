package controllers

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(c *gin.Context) {

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

func GetCourseById(c *gin.Context) {

	id := c.Param("id")
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

func UpdateCourseById(c *gin.Context) {
	id := c.Param("id")
	var student models.Course

	var Body struct {
		Name string
		Age  int
	}

	c.Bind(&Body)

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&student, "id = ?", id)
	initializers.DB.Model(&student).Where("id = ?", id).Update(Body.Name, Body.Age)

}

func DeleteCourseById(c *gin.Context) {
	id := c.Param("id")
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
