package controller

import (
	"github.com/gin-gonic/gin"
)

func initializeAPIGroup(r *gin.RouterGroup) {
	// Accounts
	r.GET("/accounts", GetAccounts)
	r.GET("/accounts/:id", GetAccountById)
	r.POST("/accounts", CreateNewAccount)
	r.DELETE("/accounts/:id", DeleteAccount)

	// Auth
	// TODO:

}

func InitializeHandlers(r *gin.Engine) {
	initializeAPIGroup(r.Group("/v1"))

	r.Run()
}
