package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const portNumber = ":8080"

// main is the main function
func main() {

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	r := gin.Default()

	r.Run()
}
