package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"os"
	"fmt"
	"./model"
	"./repo"
	"log"
)

func main() {

	args := os.Args[1:]

	fmt.Println(args)

	var storyId string = args[0]

	log.Printf("StoryId: [%s]", storyId)

	// Load the Story
	var story *model.Story = repo.RepoFindStory(storyId)

	log.Printf("StoryTitle: %s", story.Title)

	// Check it has NamedEntities
	if story.NamedEntities.IsEmpty() {
		log.Printf("Story [%s] has no NamedEntities", storyId)
		os.Exit(0)
	}

	// Analyse Entities (ie, call sentiment-api)
	Analyse(story.NamedEntities)

	// Load list of matched companies
	for _, companyId := range story.MatchedCompanies {
		log.Printf("Checking Company: [%s]", companyId)
		var company *model.Company = repo.RepoFindCompany(companyId)

		log.Printf("Checking Company: [%s] Name [%s]", company, company.Name)
	}

	// For each company
	//		Find entities from story that match company
	//		Construct StorySentiment on matches
	//		Count the occurrences of that name in the company information
	//		Apply that count as a multiplier on the sentiment
	//		Save the StorySentiment


	// For each company with an updated sentiment send the Id on Queue SentimentUpdated

}