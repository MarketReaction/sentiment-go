package model

import "time"

type EntitySentiment struct {
	Entity    string `json:"entity"`
	Sentiment int    `json:"sentiment"`
}
