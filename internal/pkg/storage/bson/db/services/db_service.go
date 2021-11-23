package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	VideoCollection      = "videos"
	CategoriesCollection = "categories"
)

type DatabaseService struct {
	*mongo.Database
}

func ProvideDatabaseService() DatabaseService {
	server := mountServerConnection(os.Getenv("ENV"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_NAME"))

	clientOptions := options.Client().ApplyURI(server)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return DatabaseService{}
	}

	return DatabaseService{client.Database(os.Getenv("APP_DB_NAME"))}
}

func mountServerConnection(env, user, password, hostname, dbname string) string {
	if env == "dev" || env == "" {
		return "mongodb://mongo:27017/dev_env"
	}
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, hostname, dbname)
}

func makeFindOptions(filter string, page int64, pageSize int64) (bson.M, *options.FindOptions) {
	collectionFilter := bson.M{}
	findOptions := options.Find()
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip((page - 1) * pageSize)
	if filter != "" {
		collectionFilter = bson.M{"titulo": bson.M{"$regex": fmt.Sprintf(".*%s.*", filter)}}
	}
	return collectionFilter, findOptions
}
