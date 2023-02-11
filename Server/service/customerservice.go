package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/repository"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

func GetClientSecret(c *gin.Context) {
	fmt.Println("Received GetClientSecret Request")
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	type Requestbody struct {
		Name    string
		Bill    int
		Courses string
	}

	var requestbody Requestbody

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Below is the string format of payload")
		fmt.Println(string(p))

		jsonErr := json.Unmarshal([]byte(string(p)), &requestbody)
		if jsonErr != nil {
			log.Println(jsonErr)
		}
		fmt.Println(requestbody.Bill)

		// Create a PaymentIntent with amount and currency
		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(int64(requestbody.Bill * 100)),
			Currency: stripe.String(string(stripe.CurrencyINR)),
			AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
				Enabled: stripe.Bool(true),
			},
		}

		pi, err := paymentintent.New(params)
		log.Printf("pi.New: %v", pi.ClientSecret)

		if err != nil {
			log.Printf("pi.New: %v", err)
			return
		}

		ws.WriteMessage(messageType, []byte(pi.ClientSecret))
	}
}

func MakePayment(c *gin.Context) {

	type Requestbody struct {
		Name    string
		Bill    int
		Courses string
	}
	var requestbody Requestbody
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
