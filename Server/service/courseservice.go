package service

import (
	"server/repository"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(c *gin.Context) {
	repository.GetAllCoursesFromDB(c)
}

func GetCourseById(c *gin.Context) {
	id := c.Param("id")
	repository.GetCourseByIdFromDB(id, c)
}

func UpdateCourseById(c *gin.Context) {
	id := c.Param("id")
	repository.UpdateCourseInDBById(id, c)
}

func DeleteCourseById(c *gin.Context) {
	id := c.Param("id")
	repository.DeleteCourseById(id, c)
}
