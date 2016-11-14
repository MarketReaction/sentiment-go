package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"os"
	"fmt"
	"./model"
	"./repo"
)

func main() {

	args := os.Args[1:]

	fmt.Println(args)

	var storyId string = args[0]

	fmt.Println("StoryId: " + storyId)

	// Load the Story
	var story *model.Story = repo.RepoFindStory(storyId)

	fmt.Println("StoryTitle: " + story.Title)

	// Check it has NamedEntities
	//if story.NamedEntities() == nil {
	//	log.Printf("Story [%s] has no NamedEntities", storyId)
	//	os.Exit(0)
	//}

	// Analyse Entities (ie, call sentiment-api)

	// Load list of matched companies

	// For each company
	//		Find entities from story that match company
	//		Construct StorySentiment on matches
	//		Count the occurrences of that name in the company information
	//		Apply that count as a multiplier on the sentiment
	//		Save the StorySentiment


	// For each company with an updated sentiment send the Id on Queue SentimentUpdated

}