package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ConnectionString = "mongodb://localhost:27017/"
	DatabaseName     = "ExpenseTracker"
)

var TotalCollection *mongo.Collection
var UserCollection *mongo.Collection
var TransactionCollection *mongo.Collection

func InitDatabase() {
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database(DatabaseName).Collection("User")
	TotalCollection = client.Database(DatabaseName).Collection("Amount")
	TransactionCollection = client.Database(DatabaseName).Collection("Transaction")
	fmt.Println("Database Connected successfully")
}
