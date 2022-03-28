package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderProduct struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string  `json:"name,omitempty" bson:"name,omitempty"`
	Image            string  `json:"image,omitempty" bson:"image,omitempty"`
	Price            float64 `json:"price,omitempty" bson:"price,omitempty"`
}

type Order struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	OrderProduct     []OrderProduct     `json:"order_product,omitempty" bson:"order,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
}
