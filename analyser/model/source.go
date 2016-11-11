package model

import (
	"labix.org/v2/mgo/bson"
)

type Source struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	name          string        `json:"name"`
	exclusionList []string        `json:"exclusionList"`
}

type Sources []Source
