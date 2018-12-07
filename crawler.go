package main

import (
    "github.com/mongodb/mongo-go-driver/mongo"
    "context"
    "time"
    "log"
)

func main() {
    client, err := mongo.NewClient("mongodb://localhost:27017")
    if err != nil {
        log.Fatal("Could not connect: ", err)
        return
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        return
    }

    collection := client.Database("test_database").Collection("test_collection")

    count, err := collection.Count(ctx, nil)

    log.Println("Count is: ", count)
}
