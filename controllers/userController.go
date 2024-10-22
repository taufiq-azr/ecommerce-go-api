package controllers

import (
	"context"
	"net/http"
	"time"
	"github.com/taufiq-azr/ecommerce-go-api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var UserCollection *mongo.Collection

// Fungsi untuk mengatur koleksi User
func SetupUserController(Collection *mongo.Collection) {
	UserCollection = Collection
}

// CreateUser menambahkan pengguna baru
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Set user details
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Insert into MongoDB
	_, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetUsers mendapatkan semua pengguna
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	cursor, err := UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to decode users",
			"error":   err.Error(),
		})
	}

	return c.JSON(users)
}

// UpdateUser memperbarui pengguna berdasarkan ID
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Failed to fetch user",
			"error":   err.Error(),
		})
	}

	user.UpdatedAt = time.Now()

	_, err := UserCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.JSON(user)
}

// DeleteUser menghapus pengguna berdasarkan ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := UserCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
