package config

import (
	"context"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/shinhagunn/shop-product/config/collection"
	"github.com/shinhagunn/shop-product/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mgm.SetDefaultConfig(nil, "productDB", options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))

	log.Println("Connected to ProductDB!")

	collection.Product = mgm.Coll(&models.Product{})
	collection.Category = mgm.Coll(&models.Category{})
	collection.Slide = mgm.Coll(&models.Slide{})
	collection.User = mgm.Coll(&models.User{})
	collection.Order = mgm.Coll(&models.Order{})
	collection.Comment = mgm.Coll(&models.Comment{})

	collection.User.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"uid": 1},
			Options: options.Index().SetUnique(true),
		},
	})

	collection.Category.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	})
}
