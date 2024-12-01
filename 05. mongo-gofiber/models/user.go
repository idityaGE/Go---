package models

import (
	"github.com/idityaGE/go-mongo-gofiber/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email"`
	Age      int16              `json:"age"`
	IsActive bool               `json:"is_active"`
	Address  Address            `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

var UserCol *mongo.Collection

func init() {
	database.Connect()
	UserCol = database.MG.DB.Collection("user")
}