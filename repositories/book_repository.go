package repositories

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    //"example.com/go-mongo-app/models"
    "log"
    "os"  

)

type BookRepository struct {
    collection *mongo.Collection
}

func NewBookRepository() *BookRepository {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("library").Collection("books")
    return &BookRepository{collection}
}

func (r *BookRepository) RemoveBookByID(id string) (bool, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return false, err
    }

    result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        return false, err
    }

    return result.DeletedCount > 0, nil
}
