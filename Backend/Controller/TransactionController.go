package controller

import (
	helper "ExpenseTacker/Helper"
	model "ExpenseTacker/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostTransactionController(c *gin.Context) {
	var transaction model.TransactionModel

	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in binding json": err.Error()})
		return
	}
	UserId, err := helper.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in getting user id": err.Error()})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in converting id": err.Error()})
		return
	}

	transaction.User_ID = userObjectID

	err = helper.TransactionHelperPost(transaction)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in posting transaction": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
func GetTransactionsController(c *gin.Context) {
	UserId, err := helper.GetUserIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in getting user id": err.Error()})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	transaction, err := helper.TransactionHelperGet(userObjectID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error in getting transaction": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
