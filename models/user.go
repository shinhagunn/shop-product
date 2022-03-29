package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartProduct struct {
	ProductID primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Quantity  int64              `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type UserProfile struct {
	Fullname string `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Age      string `json:"age,omitempty" bson:"age,omitempty"`
	Address  string `json:"address,omitempty" bson:"address,omitempty"`
	Gender   string `json:"gender,omitempty" bson:"gender,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UID              string        `json:"uid,omitempty" bson:"uid,omitempty"`
	Email            string        `json:"email,omitempty" bson:"email,omitempty"`
	Role             string        `json:"role,omitempty" bson:"role,omitempty"`
	UserProfile      UserProfile   `json:"user_profile,omitempty" bson:"user_profile,omitempty"`
	Cart             []CartProduct `json:"cart,omitempty" bson:"cart,omitempty"`
}
