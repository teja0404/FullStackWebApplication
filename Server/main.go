package main

import (
	"fmt"
	"server/controllers"
	"server/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("Hello, Modules!")
	initializers.PrintHello()
	initializers.LoadEnvVariables()
	initializers.EstablishConnection()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/ping", controllers.PostRequest)
	r.GET("/getall", controllers.GetAll)
	r.GET("/get/:id", controllers.GetByIndex)
	r.PUT("/update/:id", controllers.UpdateByIndex)
	r.DELETE("/delete/:id", controllers.DeleteByIndex)

	r.Run()

}
