package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Total struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	User_ID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	TotalAmount int64              `json:"total_amount"`
}

// 2nd model holding the total amount of the user
