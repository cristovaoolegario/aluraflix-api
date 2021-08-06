package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DatabaseService struct {
	Server string
	Database string
	db *mongo.Database
}

var videosCollection *mongo.Collection
var categoriesCollection *mongo.Collection

const(
	VideoCollection = "videos"
	CategoriesCollection = "categories"
)

func (dbService *DatabaseService) Connect() {
	clientOptions := options.Client().
		ApplyURI(dbService.Server)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	dbService.db = client.Database(dbService.Database)
	videosCollection = dbService.db.Collection(VideoCollection)
	categoriesCollection = dbService.db.Collection(CategoriesCollection)
}