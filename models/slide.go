package models

import "github.com/kamva/mgm/v3"

type Slide struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title,omitempty" bson:"title,omitempty"`
	Description      string `json:"description,omitempty" bson:"description,omitempty"`
	Image            string `json:"image,omitempty" bson:"image,omitempty"`
}
