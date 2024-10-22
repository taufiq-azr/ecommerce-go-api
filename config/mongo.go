package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"fmt"
)

var DB *mongo.Database

func ConnectDB() {
	
	clientOption := options.Client().ApplyURI("")

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
	
	fmt.Println("Connected to MongoDB!")

}