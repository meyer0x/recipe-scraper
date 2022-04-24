package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Recipe struct {
	Name          string   `bson:"name" json:"name"`
	Ingredients   []string `bson:"ingredients" json:"ingredients"`
	ServingPerson int      `bson:"servingPerson" json:"servingPerson"`
	URL           string   `bson:"url" json:"url"`
}

var db *mongo.Client

var recipesCollections *mongo.Collection

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Client connected MongoDB")

	db = client

	recipesCollections = client.Database("recipe-scraper").Collection("recipes")
}

func InsertManyRecipes(recipes []Recipe) {
	var recipesToInsert []interface{}
	for _, t := range recipes {
		if !isDuplicate(t) {
			recipesToInsert = append(recipesToInsert, t)
		}
	}
	recipesCollections.InsertMany(context.TODO(), recipesToInsert)
	fmt.Println("Recipes inserted", len(recipesToInsert))
}

func isDuplicate(recipe Recipe) bool {
	filter := bson.M{
		"url": recipe.URL,
	}

	cursor, err := recipesCollections.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	var recipesFiltered []bson.M

	if err = cursor.All(context.TODO(), &recipesFiltered); err != nil {
		panic(err)
	}

	if len(recipesFiltered) > 0 {
		return true
	}

	return false
}
