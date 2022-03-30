package models

import (
	"github.com/kamva/mgm/v3"
)

type Custommer struct {
	mgm.DefaultModel `bson:",inline"`
	Fullname         string `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Email            string `json:"email,omitempty" bson:"email,omitempty"`
	Phone            string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address          string `json:"address,omitempty" bson:"address,omitempty"`
	Message          string `json:"message,omitempty" bson:"message,omitempty"`
}
