package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"fmt"
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"github.com/MarketReaction/sentiment-go/analyser/repo"
	"github.com/streadway/amqp"
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

					storySentiment := &model.StorySentiment{
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

		if (companyUpdated) {

			var activeMQUrl string = fmt.Sprintf("amqp://%s:%s/", os.Getenv("ACTIVEMQ_PORT_61616_TCP_ADDR"), os.Getenv("ACTIVEMQ_PORT_61616_TCP_PORT"))

			log.Printf("ActiveMQ at url [%s]", activeMQUrl)

			conn, err := amqp.Dial(activeMQUrl)
			if err != nil {
				log.Fatalf(err, "Failed to connect to ActiveMQ")
			}

			defer conn.Close()

			//ch, err := conn.Channel()
			//log.Println(err, "Failed to open a channel")
			//defer ch.Close()
			//
			//q, err := ch.QueueDeclare(
			//	"SentimentUpdated", // name
			//	false,   // durable
			//	false,   // delete when unused
			//	false,   // exclusive
			//	false,   // no-wait
			//	nil,     // arguments
			//)
			//log.Println(err, "Failed to declare a queue")
			//
			//err = ch.Publish(
			//	"",     // exchange
			//	q.Name, // routing key
			//	false,  // mandatory
			//	false,  // immediate
			//	amqp.Publishing {
			//		Body:        []byte(company.Id.Hex()),
			//	})
			//log.Println(err, "Failed to publish a message")

		}

	}

	// For each company
	//		Find entities from story that match company
	//		Construct StorySentiment on matches
	//		Count the occurrences of that name in the company information
	//		Apply that count as a multiplier on the sentiment
	//		Save the StorySentiment

	// For each company with an updated sentiment send the Id on Queue SentimentUpdated

}
