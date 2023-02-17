package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"server/repository"

	"github.com/tidwall/gjson"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72/webhook"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "Fullstack Application BE is running",
	})

}

type Requestbody struct {
	Name           string
	InstructorName string
	Price          int
	Description    string
	Duration       int
}

func AddCourse(c *gin.Context) {
	var requestbody Requestbody
	c.Bind(&requestbody)
	repository.AddCourseInDB(requestbody.Name, requestbody.InstructorName, requestbody.Price, requestbody.Description, requestbody.Duration, c)
}

func HandleStripeWebhook(c *gin.Context) {
	fmt.Println("Inside HandleStripeWebhook method")
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sigHeader := c.GetHeader("Stripe-Signature")
	if sigHeader == "" {
		fmt.Println("No signature header found")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(os.Getenv("STRIPE_WEBHOOK_SECRET"))

	// Verify the webhook signature using your Stripe webhook secret
	event, err := webhook.ConstructEvent(payload, sigHeader, os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		fmt.Println("Error verifying webhook signature:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Handle the webhook event based on its type
	switch event.Type {
	case "payment_intent.succeeded":
		fmt.Println("Received payment intent succeeded notification from Stripe")
		value := gjson.Get(string(payload), "data.object.client_secret")
		repository.UpdatePaymentToSuccess(value.String())

	case "payment_intent.failed":
		fmt.Println("Received payment intent failed notification from Stripe")
		value := gjson.Get(string(payload), "data.object.client_secret")
		repository.UpdatePaymentToSuccess(value.String())

	default:
		fmt.Println("Unhandled event type:", event.Type)
	}

	c.Status(http.StatusOK)
}
