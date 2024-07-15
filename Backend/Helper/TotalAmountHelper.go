package helper

import (
	// database "ExpenseTracker/database"
	// model "ExpenseTracker/model"
	database "ExpenseTacker/Database"
	model "ExpenseTacker/Model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostTotalAmountHelper(totalAmount model.Total) error {
	totalAmount.ID = primitive.NewObjectID()

	_, err := database.Collection.InsertOne(context.Background(), totalAmount)
	if err != nil {
		return fmt.Errorf("error in inserting total amount: %v", err)
	}
	return nil
}

func GetTotalAmountHelper(userID primitive.ObjectID) ([]model.Total, error) {
	var totals []model.Total

	cursor, err := database.Collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("error in finding totals: %v", err)
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &totals); err != nil {
		return nil, fmt.Errorf("error in decoding totals: %v", err)
	}

	return totals, nil
}
