package controller

import (
	"github.com/gin-gonic/gin"
)

func initializeAPIGroup(r *gin.RouterGroup) {
	// Accounts
	r.GET("/accounts", getAccounts)
	r.GET("/accounts/:id", getAccountById)
	r.POST("/accounts", createNewAccount)
	r.DELETE("/accounts/:id", deleteAccount)
}

func InitializeHandlers(r *gin.Engine) {
	initializeAPIGroup(r.Group("/v1"))

	r.Run()
}
