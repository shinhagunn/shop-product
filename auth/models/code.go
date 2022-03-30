package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Code struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Code             string             `json:"code,omitempty" bson:"code,omitempty"`
	CodeExpiration   time.Time          `json:"codeExpiration,omitempty" bson:"codeExpiration,omitempty"`
	State            string             `json:"state,omitempty" bson:"state,omitempty"`
}
