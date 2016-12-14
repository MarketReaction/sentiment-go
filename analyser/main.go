package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"fmt"
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"github.com/MarketReaction/sentiment-go/analyser/repo"
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
		log.Printf("Checking Company: [%s]", companyId)
		var company *model.Company = repo.RepoFindCompany(companyId)

		log.Printf("Checking Company: [%s] Name [%s]", company.Id, company.Name)

		for _, storyOrg := range story.NamedEntities.Organisations {

			for _, companyOrg := range company.NamedEntities.Organisations {
				if storyOrg.Name == companyOrg.Name {

					log.Printf("Matched name [%s] in company [%s]", storyOrg.Name, company.Name)

					var sentimentSum int = 0

					for _, sentiment := range companyOrg.Sentiments {
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

					//final StorySentiment storySentiment = new StorySentiment(company.getId(), story.getDatePublished(), story.getId());
					//
					//int multiplier = Stream.of(company.getEntities().getOrganisations(), company.getEntities().getPeople(), company.getEntities().getLocations(), company.getEntities().getMisc()).flatMap(Collection::stream)
					//        .filter(companyNamedEntity -> companyNamedEntity.equals(namedEntity)).collect(Collectors.summingInt(NamedEntity::getCount));
					//
					//storySentiment.getEntitySentiment().add(new EntitySentiment(namedEntity.getName(), namedEntity.getSentiments().stream().collect(Collectors.summingInt(Sentiment::getSentiment)) * multiplier));
					//
					//storySentimentRepository.save(storySentiment);
					//
					//updatedCompanyIds.add(company.getId());

				}
			}
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
