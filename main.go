package main

import (
	"context"
	"log"
	"net/http"
	"time"

	sw "github.com/fsschmitt/its-your-birthday/src/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	log.Printf("Server started")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
