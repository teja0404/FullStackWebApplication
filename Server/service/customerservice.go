package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/bo"
	"server/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

// Upgrade the connection to web socket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func AddCustomer(c *gin.Context) {
	var customer bo.Customer

	c.Bind(&customer)
	repository.AddCustomerInDB(customer.Name, customer.Age, customer.Email, customer.Gender, c)
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

func GetClientSecret(c *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var payment bo.Payment

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		jsonErr := json.Unmarshal([]byte(string(p)), &payment)
		if jsonErr != nil {
			log.Println(jsonErr)
		}
		fmt.Println(payment.Bill)

		// Create a PaymentIntent with amount and currency
		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(int64(payment.Bill * 100)),
			Currency: stripe.String(string(stripe.CurrencyINR)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}

		pi, err := paymentintent.New(params)

		if err != nil {
			log.Printf("pi.New: %v", err)
			return
		}

		ws.WriteMessage(messageType, []byte(pi.ClientSecret))
		initiatePayment(payment.Name, payment.Courses, payment.Bill, "Processing", pi.ClientSecret)
	}
}

func initiatePayment(Name string, Courses string, Bill int, Status string, ClientSecret string) {
	dt := time.Now()
	var date = dt.Format(time.UnixDate)
	repository.InitiatePayment(Name, Courses, Bill, date, Status, ClientSecret)
}

func MakePayment(c *gin.Context) {
	var payment bo.Payment

	dt := time.Now()
	//Format date
	var date = dt.Format(time.UnixDate)
	c.Bind(&payment)
	repository.PersistPaymentInDB(payment.Name, payment.Courses, payment.Bill, date, c)
}

func GetPaymentsByName(c *gin.Context) {
	name := c.Param("name")
	repository.GetPaymentsByName(name, c)
}
