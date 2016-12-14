package model

type EntitySentiment struct {
	Entity    string `json:"entity"`
	Sentiment int    `json:"sentiment"`
}
