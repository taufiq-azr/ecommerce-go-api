package routes

import (
	"github.com/taufiq-azr/ecommerce-go-api/controllers"
	"github.com/taufiq-azr/ecommerce-go-api/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // Mengelompokkan route di bawah /api

	// Route untuk kategori
	api.Post("/categories", controllers.CreateCategory)       // Menambahkan kategori
	api.Put("/categories/:id", controllers.UpdateCategory)    // Memperbarui kategori
	api.Delete("/categories/:id", controllers.DeleteCategory) // Menghapus kategori
	api.Get("/categories", controllers.GetCategories)         // Mendapatkan semua kategori

	// Route untuk produk
	api.Post("/products", middleware.AuthMiddleware, controllers.CreateProduct)       // Menambahkan produk
	api.Get("/products", controllers.GetProduct)                                     // Mendapatkan semua produk
	api.Get("/products/:id", controllers.GetProduct)                                  // Mendapatkan produk berdasarkan ID
	api.Put("/products/:id", middleware.AuthMiddleware, controllers.UpdateProduct)    // Memperbarui produk
	api.Delete("/products/:id", middleware.AuthMiddleware, controllers.DeleteProduct) // Menghapus produk

	api.Get("/products/category/:category_id", controllers.GetProductByIDCategory) // Mendapatkan produk berdasarkan kategori
	api.Get("/products/category/name/:category_name", controllers.GetProductsByCategoryName)


	// Route untuk user
	api.Post("/users", controllers.CreateUser)                                  // Menambahkan user
	api.Get("/users", controllers.GetUsers)                                     // Mendapatkan semua user
	api.Put("/users/:id", middleware.AuthMiddleware, controllers.UpdateUser)    // Memperbarui user
	api.Delete("/users/:id", middleware.AuthMiddleware, controllers.DeleteUser) // Menghapus user
}
