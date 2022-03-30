package config

import (
	"context"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mgm.SetDefaultConfig(nil, "authDB", options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))

	log.Println("Connected to authDB!")

	collection.Code = mgm.Coll(&models.Code{})
	collection.User = mgm.Coll(&models.User{})
	collection.Custommer = mgm.Coll(&models.Custommer{})

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
}
