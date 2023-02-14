package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"kordian.net/account/internal/controller"
	"kordian.net/account/internal/service/db"
)

func main() {
	fmt.Println("Hello World!")

	db.InitializeDatabase()
	controller.InitializeHandlers(gin.Default())
}
