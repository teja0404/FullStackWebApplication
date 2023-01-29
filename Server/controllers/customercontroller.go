package controllers

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AddCustomer(c *gin.Context) {

	var requestbody struct {
		Name   string
		Age    int
		Email  string
		Gender string
	}

	c.Bind(&requestbody)

	customer := models.Customer{Name: requestbody.Name, Age: requestbody.Age, Email: requestbody.Email, Gender: requestbody.Gender}

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	result := initializers.DB.Create(&customer)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func GetAllCustomers(c *gin.Context) {

	var allcustomers []models.Customer

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Find(&allcustomers)

	c.JSON(http.StatusOK, gin.H{
		"allcustomers": allcustomers,
	})
}

func GetCustomerById(c *gin.Context) {

	id := c.Param("id")
	var customer models.Customer

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&customer, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func UpdateCustomerById(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	var Body struct {
		Name      string
		Age       int
		Email     string
		Gender    string
		IsDeleted bool
	}

	c.Bind(&Body)

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.First(&customer, "id = ?", id)
	//Update the fetched record and persist back
	customer.Name = Body.Name
	customer.Age = Body.Age
	customer.Email = Body.Email
	customer.Gender = Body.Gender
	customer.IsDeleted = Body.IsDeleted

	initializers.DB.Save(&customer)

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})

}

func DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Where("id = ?", id).Delete(&customer)
}

func MakePayment(c *gin.Context) {

	var requestbody struct {
		Name    string
		Bill    int
		Courses string
	}

	dt := time.Now()

	//Format date
	var date = dt.Format(time.UnixDate)

	c.Bind(&requestbody)

	payment := models.Payment{Name: requestbody.Name, Courses: requestbody.Courses, Bill: requestbody.Bill, Date: date}

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	result := initializers.DB.Create(&payment)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": payment,
	})
}

func GetPaymentsByName(c *gin.Context) {

	name := c.Param("name")
	var payments []models.Payment

	if initializers.DB == nil {
		initializers.DB = initializers.EstablishConnection()
	}

	if initializers.DB == nil {
		fmt.Println("DB Connection is not established")
		return
	}

	initializers.DB.Find(&payments, "name = ?", name)

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
	})
}
