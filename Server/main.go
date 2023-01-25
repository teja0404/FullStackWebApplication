package main

import (
	"Server/initializers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello, Modules!")
	initializers.PrintHello()
	initializers.LoadEnvVariables()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
