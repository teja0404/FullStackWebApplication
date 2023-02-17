package repository

import (
	"fmt"
	"net/http"
	"server/initializers"
	"server/models"

	"github.com/gin-gonic/gin"
)

func AddCustomerInDB(Name string, Age int, Email string, Gender string, c *gin.Context) {
	customer := models.Customer{Name: Name, Age: Age, Email: Email, Gender: Gender}

	result := initializers.DB.Create(&customer)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func GetAllCustomersFromDB(c *gin.Context) {
	var allcustomers []models.Customer

	initializers.DB.Find(&allcustomers)

	c.JSON(http.StatusOK, gin.H{
		"allcustomers": allcustomers,
	})
}

func GetCustomerByID(id string, c *gin.Context) {
	var customer models.Customer

	initializers.DB.First(&customer, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func UpdateCustomerById(id string, c *gin.Context) {
	var customer models.Customer

	var Body struct {
		Name      string
		Age       int
		Email     string
		Gender    string
		IsDeleted bool
	}

	c.Bind(&Body)

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

func UpdatePaymentToSuccess(clientsecret string) {
	fmt.Println("Updating payment to success")
	var payment models.Payment

	initializers.DB.First(&payment, "client_secret = ?", clientsecret)
	payment.Status = "Success"
	initializers.DB.Save(&payment)

}

func UpdatePaymentToFailed(clientsecret string) {
	fmt.Println("Updating payment to Failed")
	var payment models.Payment

	initializers.DB.First(&payment, "client_secret = ?", clientsecret)
	payment.Status = "Failed"
	initializers.DB.Save(&payment)

}

func DeleteCustomerById(id string, c *gin.Context) {
	var customer models.Customer

	initializers.DB.Where("id = ?", id).Delete(&customer)
}

func PersistPaymentInDB(Name string, Courses string, Bill int, Date string, c *gin.Context) {
	payment := models.Payment{Name: Name, Courses: Courses, Bill: Bill, Date: Date}

	result := initializers.DB.Create(&payment)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payment": payment,
	})
}

func InitiatePayment(Name string, Courses string, Bill int, Date string, Status string, ClientSecret string) {
	payment := models.Payment{Name: Name, Courses: Courses, Bill: Bill, Date: Date, Status: Status, ClientSecret: ClientSecret}

	result := initializers.DB.Create(&payment)

	if result.Error != nil {
		fmt.Println("Error while initiating payment")
		return
	} else {
		fmt.Println("Payment succesfully persisted")
	}

}

func GetPaymentsByName(name string, c *gin.Context) {
	var payments []models.Payment

	initializers.DB.Where("name = ? AND status = ?", name, "Success").Find(&payments)

	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
	})
}
