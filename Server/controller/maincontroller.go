package controller

//This is a Controller class
import (
	"server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	//Get the default Gin router
	r := gin.Default()

	// cors.Default() setup the middleware with default options.
	// Accepts requests (GET, POST, PUT, DELETE) from all origins
	r.Use(cors.Default())

	//Health check
	//To check whether application is running or not
	r.GET("/healthcheck", service.HealthCheck)

	//Controllers related to Customers
	r.POST("/addcustomer", service.AddCustomer)
	r.GET("/getallcustomers", service.GetAllCustomers)
	r.GET("/getcustomerbyid/:id", service.GetCustomerById)
	r.PUT("/updatecustomerbyid/:id", service.UpdateCustomerById)
	r.DELETE("/deletecustomerbyid/:id", service.DeleteCustomerById)
	r.GET("/makepayment", service.MakePayment)
	r.GET("/getpaymentsbyname/:name", service.GetPaymentsByName)

	//Controllers related to Course
	//Only admin can add the Course
	r.POST("/addcourse", service.AddCourse)
	r.GET("/getallcourses", service.GetAllCourses)
	r.GET("/getcoursebyid/:id", service.GetCourseById)
	r.PUT("/updatecoursebyid/:id", service.UpdateCourseById)
	r.DELETE("/deletecoursebyid/:id", service.DeleteCourseById)

	//Runs on the Port which we configured in Env file
	r.Run()

}
