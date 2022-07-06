package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//func DatabaseInstance() *mongo.Client {
//mongoDB := "mongodb://localhost:27017"
//fmt.Print(mongoDB)
//client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
//if err != nil {
//	log.Fatalln(err)
//}
//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//defer cancel()
//
//err = client.Connect(ctx)
//if err != nil {
//	log.Fatal(err)
//}
//
//fmt.Println("\nconnected to mongodb")
//
//return client
//}

//var Client *mongo.Client = DatabaseInstance()

// database open collection
//func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
//	var collection *mongo.Collection = client.Database("mongodb1").Collection("posts")
//	return collection
//}

// *******************************************************//

func DatabaseInstance() *mongo.Client {
	mongoDB := "mongodb://localhost:27017"

	fmt.Print(mongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDB))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nconnected to mongodb")

	return client
}

var Client *mongo.Client = DatabaseInstance()

// database open collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("mongodb1").Collection(collectionName)
	return collection
}
