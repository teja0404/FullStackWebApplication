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
	r.POST("/webhooks/stripe", service.HandleStripeWebhook)

	//Controllers related to Customers
	r.POST("/customer", service.AddCustomer)
	r.GET("/customers", service.GetAllCustomers)
	r.GET("/customer/:id", service.GetCustomerById)
	r.PUT("/customer/:id", service.UpdateCustomerById)
	r.DELETE("/customer/:id", service.DeleteCustomerById)
	r.GET("/clientsecret", service.GetClientSecret)
	r.POST("/payment", service.MakePayment)
	r.GET("/payment/:name", service.GetPaymentsByName)

	//Controllers related to Course
	//Only admin can add the Course
	r.POST("/course", service.AddCourse)
	r.GET("/courses", service.GetAllCourses)
	r.GET("/course/:id", service.GetCourseById)
	r.PUT("/course/:id", service.UpdateCourseById)
	r.DELETE("/course/:id", service.DeleteCourseById)

	//Runs on the Port which we configured in Env file
	r.Run()
}
