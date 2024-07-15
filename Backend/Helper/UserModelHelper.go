package helper

import (
	database "ExpenseTacker/Database"
	model "ExpenseTacker/Model"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("your_secret_key")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func RegisterUserHelper(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error in hashing password: %v", err)
	}
	user.Password = string(hashedPassword)

	user.ID = primitive.NewObjectID()

	_, err = database.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error in inserting user: %v", err)
	}
	return nil
}

func LoginUserHelper(user model.User) (string, error) {
	var foundUser model.User
	err := database.Collection.FindOne(context.Background(), bson.M{"name": user.Name}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: foundUser.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
	})
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUserIDFromContext(c *gin.Context) (string, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", errors.New("user ID not found in context")
	}

	userIDString, ok := userID.(string)
	if !ok {
		return "", errors.New("user ID is not a string")
	}

	return userIDString, nil
}

func GetAllRegisterUser() ([]model.User, error) {
	var users []model.User

	cursor, err := database.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error in finding users: %v", err)
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("error in decoding users: %v", err)
	}

	return users, nil
}
