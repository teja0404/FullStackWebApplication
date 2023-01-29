package main

import (
	"server/controllers"
	"server/initializers"
	"server/migration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializers.PrintHello()
	initializers.LoadEnvVariables()
	migration.MigrateDatabases()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	//Health check
	r.GET("/healthcheck", controllers.HealthCheck)

	//Controllers related to Customers
	r.POST("/addcustomer", controllers.AddCustomer)
	r.GET("/getallcustomers", controllers.GetAllCustomers)
	r.GET("/getcustomerbyid/:id", controllers.GetCustomerById)
	r.PUT("/updatecustomerbyid/:id", controllers.UpdateCustomerById)
	r.DELETE("/deletecustomerbyid/:id", controllers.DeleteCustomerById)
	r.POST("/makepayment", controllers.MakePayment)
	r.GET("/getpaymentsbyname/:name", controllers.GetPaymentsByName)

	//Controllers related to Course
	//Only admin can add the Course
	r.POST("/addcourse", controllers.AddCourse)

	r.GET("/getallcourses", controllers.GetAllCourses)
	r.GET("/getcoursebyid/:id", controllers.GetCourseById)
	r.PUT("/updatecoursebyid/:id", controllers.UpdateCourseById)
	r.DELETE("/deletecoursebyid/:id", controllers.DeleteCourseById)

	r.Run()

}
