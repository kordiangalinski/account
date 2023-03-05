package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"kordian.net/account/internal/controller"
)

func main() {
	fmt.Println("Hello World!")

	controller.InitializeHandlers(gin.Default())
}
