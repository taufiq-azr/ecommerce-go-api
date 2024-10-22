package config

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/taufiq-azr/ecommerce-go-api/controllers"
)

type Collections struct {
	UserCollection     *mongo.Collection
	ProductCollection  *mongo.Collection
	CategoryCollection *mongo.Collection
	// Tambahkan koleksi lainnya sesuai kebutuhan
}

var DB Collections

func ConnectDB() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/go-ecommerce")

	// Membuat koneksi ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}

	// Cek koneksi ke database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set global database and collections
	DB.UserCollection = client.Database("go-ecommerce").Collection("users")
	DB.ProductCollection = client.Database("go-ecommerce").Collection("products")
	DB.CategoryCollection = client.Database("go-ecommerce").Collection("categories")

	controllers.SetupUserController(DB.UserCollection)
	controllers.SetupProductController(DB.ProductCollection)
	controllers.SetupCategoryController(DB.CategoryCollection)

	fmt.Println("Connected to MongoDB!")
}
