package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/idityaGE/go-mongo-gofiber/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var logger *zap.Logger

// SetLogger sets the logger instance for the package
func SetLogger(l *zap.Logger) {
	logger = l
}

// GetAllUsers retrieves all users
func GetAllUsers(c *fiber.Ctx) error {
	logger.Info("Fetching all users")
	var users []models.User
	cursor, err := models.UserCol.Find(context.TODO(), bson.D{})
	if err != nil {
		logger.Error("Failed to retrieve users", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	if err := cursor.All(context.TODO(), &users); err != nil {
		logger.Error("Failed to decode users", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to decode users",
		})
	}
	logger.Info("Successfully fetched all users", zap.Int("count", len(users)))
	return c.JSON(users)
}

// GetUserById retrieves a user by ID
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		logger.Warn("ID not provided")
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Warn("Invalid ID format", zap.String("id", id))
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var user models.User
	err = models.UserCol.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		logger.Warn("User not found", zap.String("id", id))
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	} else if err != nil {
		logger.Error("Failed to retrieve user", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	logger.Info("Successfully retrieved user", zap.String("id", id))
	return c.JSON(user)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		logger.Warn("Failed to parse request body", zap.Error(err))
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := models.UserCol.InsertOne(context.TODO(), user)
	if err != nil {
		logger.Error("Failed to create user", zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	logger.Info("User created successfully", zap.String("id", user.ID.Hex()))
	return c.JSON(user)
}

// UpdateUser updates an existing user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		logger.Warn("ID not provided")
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Warn("Invalid ID format", zap.String("id", id))
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		logger.Warn("Failed to parse request body", zap.Error(err))
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	_, err = models.UserCol.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": user})
	if err != nil {
		logger.Error("Failed to update user", zap.String("id", id), zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	logger.Info("User updated successfully", zap.String("id", id))
	return c.JSON(user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		logger.Warn("ID not provided")
		return c.Status(400).JSON(fiber.Map{
			"error": "ID is required",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Warn("Invalid ID format", zap.String("id", id))
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	_, err = models.UserCol.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		logger.Error("Failed to delete user", zap.String("id", id), zap.Error(err))
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	logger.Info("User deleted successfully", zap.String("id", id))
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
