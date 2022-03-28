package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	ProductID        primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Content          string             `json:"content,omitempty" bson:"content,omitempty"`
	LikeCount        int64              `json:"like_count,omitempty" bson:"like_count,omitempty"`
	DislikeCount     int64              `json:"dislike_count,omitempty" bson:"dislike_count,omitempty"`
}
