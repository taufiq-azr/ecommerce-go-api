package controllers

import (
	"context"
	"github.com/taufiq-azr/ecommerce-go-api/models"
	"net/http"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

func SetupProductController(collection *mongo.Collection) {
	productCollection = collection
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validasi ID kategori
	categoryID, err := primitive.ObjectIDFromHex(product.CategoryID.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid category ID",
		})
	}
	product.CategoryID = categoryID

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err = productCollection.InsertOne(context.TODO(), product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to create product",
			"error":   err.Error(),
		})
	}

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	// Validasi ID produk
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid product ID",
		})
	}

	var product models.Product
	err = productCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Failed to parse request body",
			"error":   err.Error(),
		})
	}

	// Validasi ID kategori
	categoryID, err := primitive.ObjectIDFromHex(product.CategoryID.Hex())
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid category ID",
		})
	}
	product.CategoryID = categoryID

	product.UpdatedAt = time.Now()

	_, err = productCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": product})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to update product",
			"error":   err.Error(),
		})
	}

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}

func GetProductByIDCategory(c *fiber.Ctx) error {
	categoryID := c.Params("category_id") // Mengambil category_id dari parameter URL

	// Validasi ID kategori
	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Invalid category ID",
			"error":   err.Error(),
		})
	}

	var products []models.Product
	// Mencari produk berdasarkan category_id
	cursor, err := productCollection.Find(context.TODO(), bson.M{"category_id": objectID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch products",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &products); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to decode products",
			"error":   err.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "No products found for this category",
			"error":   err.Error(),
		})
	}

	return c.JSON(products) // Mengembalikan daftar produk yang ditemukan
}

// GetProductsByCategoryName mencari produk berdasarkan nama kategori
func GetProductsByCategoryName(c *fiber.Ctx) error {
	categoryName := c.Params("category_name") // Ambil nama kategori dari parameter URL

	// Mencari ID kategori berdasarkan nama
	var category models.Category
	err := categoryCollection.FindOne(context.TODO(), bson.M{"name": categoryName}).Decode(&category)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "Category not found",
			"error":   err.Error(),
		})
	}

	// Mencari produk berdasarkan ID kategori
	var products []models.Product
	cursor, err := productCollection.Find(context.TODO(), bson.M{"category_id": category.ID.Hex()}) // Ubah ke Hex
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to fetch products",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &products); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to decode products",
			"error":   err.Error(),
		})
	}

	return c.JSON(products) // Mengembalikan daftar produk yang ditemukan
}
