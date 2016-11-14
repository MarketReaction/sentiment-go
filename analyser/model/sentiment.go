package model

type Sentiment struct {
	Sentence  string     `json:"sentence"`
	Sentiment int        `json:"sentiment"`
}
