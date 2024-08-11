package dbmongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNewMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MongodbUri"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

type MongoRepo struct {
	db         *mongo.Database
	collection string
}
func NewMongoRepository(db *mongo.Database, collection string) *MongoRepo {
	return &MongoRepo{
		db:         db,
		collection: collection,
	}
}
func NewCollection(dbname ,taskCollectionName string) *MongoRepo{
	client := GetNewMongoClient()
	return  NewMongoRepository(client.Database(dbname), taskCollectionName)

}