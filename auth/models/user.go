package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	UID              string `json:"uid,omitempty" bson:"uid,omitempty"`
	Email            string `json:"email,omitempty" bson:"email,omitempty"`
	Password         string `json:"password,omitempty" bson:"password,omitempty"`
	State            string `json:"state,omitempty" bson:"state,omitempty"`
	Role             string `json:"role,omitempty" bson:"role,omitempty"`
}
