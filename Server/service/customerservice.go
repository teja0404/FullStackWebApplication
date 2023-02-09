package service

import (
	"fmt"
	"log"
	"net/http"
	"server/repository"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func reader(conn *websocket.Conn) {
	for {

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// printing out the message for confirmation
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func MakePayment(c *gin.Context) {
	fmt.Println("This is here")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader(ws)
}

// func MakePayment(c *gin.Context) {
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer ws.Close()

// 	type Requestbody struct {
// 		Name    string
// 		Bill    int
// 		Courses string
// 	}

// 	for {
// 		var requestbody Requestbody
// 		err := ws.ReadJSON(&requestbody)
// 		if err == nil {
// 			log.Printf("error occurred: %v", err)
// 			break
// 		}
// 		log.Println(requestbody)

// 		// send message from server
// 		if err := ws.WriteJSON(requestbody); !errors.Is(err, nil) {
// 			log.Printf("error occurred: %v", err)
// 		}
// 		dt := time.Now()
// 		//Format date
// 		var date = dt.Format(time.UnixDate)
// 		c.Bind(&requestbody)
// 		repository.PersistPaymentInDB(requestbody.Name, requestbody.Courses, requestbody.Bill, date, c)
// 	}
// }

func GetPaymentsByName(c *gin.Context) {
	name := c.Param("name")
	repository.GetPaymentsByName(name, c)
}
