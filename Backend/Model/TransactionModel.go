package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// This is the main Expense app having the All expense data init
type TransactionModel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id" `
	User_ID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	Transaction string             `json:"transaction"`
	Amount      int64              `json:"amount"`
}
