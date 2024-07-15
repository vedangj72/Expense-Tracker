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
	CollectionName   = "ExpenseCollection"
)

var Collection *mongo.Collection

func InitDatabase() {
	clientOptions := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database(DatabaseName).Collection(CollectionName)
	fmt.Println("Database Connected successfully")
}
