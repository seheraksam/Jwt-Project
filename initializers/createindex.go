package initializers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateIndexes() {
	collection := Client.Database("jwt-project").Collection("users")

	// Email için unique index ekleme
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},              // Email alanına index ekle
		Options: options.Index().SetUnique(true), // Uniqe olması için
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("MongoDB Index oluşturulamadı:", err)
	}

	log.Println("✅ Unique index başarıyla oluşturuldu.")
}
