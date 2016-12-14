package model

import "time"

type StorySentiment struct {
	Company         string            `json:"company"`
	StoryDate       time.Time         `json:"storyDate"`
	Story           string            `json:"story"`
	EntitySentiment []EntitySentiment `json:"entitySentiment"`
}
