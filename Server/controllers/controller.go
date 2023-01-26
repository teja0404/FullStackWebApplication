package controllers

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func PostRequest(c *gin.Context) {

	var body struct {
		Name string
		Age  int
	}

	c.Bind(&body)

	stud := models.Student{Name: body.Name, Age: body.Age}

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	result := initializers.DB.Create(&stud)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student": stud,
	})
}

func GetAll(c *gin.Context) {

	var allStudents []models.Student

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Find(&allStudents)

	c.JSON(http.StatusOK, gin.H{
		"student": allStudents,
	})
}

func GetByIndex(c *gin.Context) {

	id := c.Param("id")
	var Student models.Student

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&Student, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"student": Student,
	})
}

func UpdateByIndex(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

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

func DeleteByIndex(c *gin.Context) {
	id := c.Param("id")
	var Student models.Student

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Where("id = ?", id).Delete(&Student)
}
