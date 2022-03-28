package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartProduct struct {
	ProductID primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity  int64              `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UID              string        `json:"uid,omitempty" bson:"uid,omitempty"`
	Email            string        `json:"email,omitempty" bson:"email,omitempty"`
	Role             string        `json:"role,omitempty" bson:"role,omitempty"`
	Cart             []CartProduct `json:"cart,omitempty" bson:"cart,omitempty"`
}
