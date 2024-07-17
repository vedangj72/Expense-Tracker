package helper

import (
	database "ExpenseTacker/Database"
	model "ExpenseTacker/Model"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TransactionHelperPost(transaction model.TransactionModel) error {
	transaction.ID = primitive.NewObjectID()

	success, err := database.TransactionCollection.InsertOne(context.Background(), transaction)
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println(success)
	return nil
}

func TransactionHelperGet(userId primitive.ObjectID) ([]model.TransactionModel, error) {
	var transactions []model.TransactionModel

	// Filter to find transactions belonging to the specific user
	filter := bson.M{"user_id": userId}
	cursor, err := database.TransactionCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cursor.Close(context.Background())

	// Decode all documents into the transactions slice
	err = cursor.All(context.Background(), &transactions)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if transactions == nil {
		transactions = []model.TransactionModel{}
	}

	return transactions, nil
}
