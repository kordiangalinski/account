package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kordian.net/account/internal/account"
)

func getAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Request OK",
	})
}

func getAccountById(c *gin.Context) {
	id := c.Param("id")
	acc := account.Account{ID: id, FirstName: "Kordian"}
	c.JSON(http.StatusOK, acc)
}

func createNewAccount(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "Account created",
	})
}

func deleteAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Account Deleted",
	})

}
