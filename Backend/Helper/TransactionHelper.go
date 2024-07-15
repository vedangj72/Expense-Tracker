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

	success, err := database.Collection.InsertOne(context.Background(), transaction)
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println(success)
	return nil
}

func TransactionHelperGet(UserId primitive.ObjectID) ([]model.TransactionModel, error) {
	var transation []model.TransactionModel

	cursor, err := database.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), &transation)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return transation, nil

}
