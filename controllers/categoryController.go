package controllers

import (
	"context"
	"e-commerce-api/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var categoryCollection *mongo.Collection

func SetupCategoryController(collection *mongo.Collection) {
	categoryCollection = collection
}

func CreateCategory(c *fiber.Ctx) error {
	var categories models.Category
	if err := c.BodyParser(&categories); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "failed to parse request body",
			"error":   err.Error(),
		})
	}

	if categories.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Category name is required"})
	}

	categories.ID = primitive.NewObjectID()
	categories.CreatedAt = time.Now()
	categories.UpdatedAt = time.Now()

	_, err := categoryCollection.InsertOne(context.TODO(), categories)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch categories",
			"error":   err.Error(),
		})
	}

	return c.JSON(categories)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	cursor, err := categoryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch categories",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &categories); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to decode categories",
			"error":   err.Error(),
		})
	}

	return c.JSON(categories)
}

// UpdateCategory memperbarui kategori berdasarkan ID
func UpdateCategory(c *fiber.Ctx) error {
	var category models.Category
	id := c.Params("id")

	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch categories",
			"error":   err.Error(),
		})
	}

	category.UpdatedAt = time.Now()

	_, err := categoryCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": category})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to update categories",
			"error":   err.Error(),
		})
	}

	return c.JSON(category)
}

// DeleteCategory menghapus kategori berdasarkan ID
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := categoryCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to delete category",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}
