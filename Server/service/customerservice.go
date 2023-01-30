package service

import (
	"server/repository"
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
	repository.AddCustomerInDB(requestbody.Name, requestbody.Age, requestbody.Email, requestbody.Gender, c)
}

func GetAllCustomers(c *gin.Context) {
	repository.GetAllCustomersFromDB(c)
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	repository.GetCustomerByID(id, c)
}

func UpdateCustomerById(c *gin.Context) {
	id := c.Param("id")
	repository.UpdateCustomerById(id, c)
}

func DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")
	repository.DeleteCustomerById(id, c)
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
	repository.PersistPaymentInDB(requestbody.Name, requestbody.Courses, requestbody.Bill, date, c)
}

func GetPaymentsByName(c *gin.Context) {
	name := c.Param("name")
	repository.GetPaymentsByName(name, c)
}
