package model

type NamedEntity struct {
	Name       string       `json:"name"`
	Count      int        	`json:"count"`
	Matched    bool        	`json:"matched"`
	Sentiments []Sentiment  `json:"sentiments"`
}
