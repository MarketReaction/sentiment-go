package model

type NamedEntity struct {
	name       string       `json:"name"`
	count      int        	`json:"count"`
	matched    bool        	`json:"matched"`
	sentiments []Sentiment  `json:"sentiments"`
}
