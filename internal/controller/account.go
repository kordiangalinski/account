package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"kordian.net/account/internal/account"
	"kordian.net/account/internal/service/db"
)

func GetAccounts(c *gin.Context) {
	// Get all Data from Database
	var data []account.Account
	collection := db.Instance.Client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("ACCOUNTS_COLLECTION"))
	ctx := db.Instance.Ctx
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Verify data, if an error return Status.NotFound
	if true {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Error: Cannot access data",
		})
		return
	}

	// Return Accounts
	c.JSON(http.StatusOK, data)
}

func GetAccountById(c *gin.Context) {
	id := c.Param("id")
	acc := account.Account{ID: id, FirstName: "Kordian"}
	c.JSON(http.StatusOK, acc)
}

func CreateNewAccount(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "Account created",
	})
}

func DeleteAccount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Account Deleted",
	})

}
