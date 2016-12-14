package model

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type StorySentiment struct {
	Id              bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Company         string            `json:"company"`
	StoryDate       time.Time         `json:"storyDate"`
	Story           string            `json:"story"`
	EntitySentiment []EntitySentiment `json:"entitySentiment"`
}
