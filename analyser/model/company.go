package model

import (
	"labix.org/v2/mgo/bson"
)

type Company struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name          string        `json:"name"`
	Exchange      string        `json:"exchange"`
	NamedEntities NamedEntities `bson:"entities"`
}

type Companies []Company
