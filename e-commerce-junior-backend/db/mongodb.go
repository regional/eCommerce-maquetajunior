package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName = "majesticdb"
)

var client *mongo.Client

func InsertDocument(collectionName string, document interface{}) error {
    client, err := connectMongoDB()
    if err != nil {
        return err
    }

    collection := client.Database(DatabaseName).Collection(collectionName)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, document)
    return err
}

func UpdateDocument(collectionName string, filter interface{}, document interface{}) error {
    client, err := connectMongoDB()
    if err != nil {
        return err
    }

    collection := client.Database(DatabaseName).Collection(collectionName)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    updateDoc := bson.M{
        "$set": document,
    }

    opts := options.Update().SetUpsert(true)    
    result, err := collection.UpdateOne(ctx, filter, updateDoc, opts)
    if err != nil {
        return err
    }

    if result.MatchedCount == 0 && result.UpsertedCount == 0 {
        return fmt.Errorf("document was not inserted or updated")
    }

    fmt.Printf("Matched %v documents and upserted %v documents.\n", result.MatchedCount, result.UpsertedCount)
    return nil
}

func GetDocuments(collectionName string, filter interface{}, result interface{}) error {
    client, err := connectMongoDB()
    if err != nil {
        return err
    }

    collection := client.Database(DatabaseName).Collection(collectionName)
    findOptions := options.Find().SetSort(bson.D{{"timestamp", 1}})

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, filter, findOptions)
    if err != nil {
        return err
    }
    defer cursor.Close(ctx)

    err = cursor.All(ctx, result)
    return err
}


func connectMongoDB() (*mongo.Client, error) {
	if client != nil {
        return client, nil
    }

	mongodbUri := os.Getenv("MONGODB_URI")
    clientOptions := options.Client().ApplyURI(mongodbUri)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return nil, err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    fmt.Println("Connected to MongoDB!")
    return client, nil
}