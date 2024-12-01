package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/idityaGE/go-mongo-gofiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(c *fiber.Ctx) error {
	var data []models.User
	cursor, err := models.UserCol.Find(context.TODO(), bson.D{})
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	if err := cursor.All(context.TODO(), &data); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	return c.JSON(data)
}
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	// Convert string ID to ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var data models.User
	err = models.UserCol.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&data)
	if err == mongo.ErrNoDocuments {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(data)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	_, err := models.UserCol.InsertOne(context.TODO(), user)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
		return nil
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	_, err := models.UserCol.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
		return nil
	}

	_, err := models.UserCol.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return nil
	}
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
