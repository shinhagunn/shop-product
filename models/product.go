package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name,omitempty" bson:"name,omitempty"`
}

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	CategoryID       primitive.ObjectID `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	Price            float64            `json:"price,omitempty" bson:"price,omitempty"`
	Discount         float64            `json:"discount,omitempty" bson:"discount,omitempty"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	Image            string             `json:"image,omitempty" bson:"image,omitempty"`
}
