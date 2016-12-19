package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"fmt"
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"github.com/MarketReaction/sentiment-go/analyser/repo"
	"github.com/jjeffery/stomp"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
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
		var company *model.Company = repo.RepoFindCompany(companyId)

		var companyUpdated bool = false

		log.Printf("Checking Company: [%s] Name [%s]", company.Id.Hex(), company.Name)

		for _, storyOrg := range story.NamedEntities.Organisations {
			for _, companyOrg := range company.NamedEntities.Organisations {
				if storyOrg.Name == companyOrg.Name {

					log.Printf("Matched name [%s] in company [%s]", storyOrg.Name, company.Name)

					var sentimentSum int = 0

					for _, sentiment := range storyOrg.Sentiments {
						sentimentSum += sentiment.Sentiment
					}

					log.Printf("Saving sentiment for story published on [%s]", story.DatePublished)

					storySentiment := &model.StorySentiment{
						Id:        bson.NewObjectId(),
						Company:   company.Id.Hex(),
						StoryDate: story.DatePublished,
						Story:     story.Id.Hex(),
						EntitySentiment: []model.EntitySentiment{
							{
								Entity:    companyOrg.Name,
								Sentiment: sentimentSum * companyOrg.Count,
							},
						},
					}

					log.Println(storySentiment)

					repo.RepoInsertStorySentiment(storySentiment)

					companyUpdated = true
				}
			}
		}

		if companyUpdated {

			var activeMQUrl string = fmt.Sprintf("%s:61613", os.Getenv("ACTIVEMQ_PORT_61616_TCP_ADDR"))

			conn, err := stomp.Dial("tcp", activeMQUrl)

			if err != nil {
				println("cannot connect to server", err.Error())
				return
			}

			//conn.Send("SentimentUpdated", "", []byte(company.Id.Hex()), nil)

			err = conn.Send(
				"/queue/SentimentUpdated", // destination
				"text/plain",              // content-type
				[]byte(company.Id.Hex()))  // body
			if err != nil {
				println("cannot connect to server", err.Error())
			}

			conn.Disconnect()

			log.Printf("Message Sent for SentimentUpdated to company [%s]", company.Name)
		}
	}
}
