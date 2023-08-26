package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://root:code1234@localhost:32768"
const dbName = "netflix"
const colName = "watchlist"

// This very important
var Collection *mongo.Collection

// Connect to DB
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Connected to Collection")
}
