package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	mgm.DefaultModel `bson:",inline"`
	UserIDs          []primitive.ObjectID `json:"user_ids,omitempty" bson:"user_ids,omitempty"`
	Name             string               `json:"name,omitempty" bson:"name,omitempty"`
}

type Message struct {
	mgm.DefaultModel `bson:",inline"`
	ChatID           primitive.ObjectID `json:"chat_id,omitempty" bson:"chat_id,omitempty"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Content          string             `json:"content,omitempty" bson:"content,omitempty"`
}
