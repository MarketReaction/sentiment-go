package repo

import (
	"log"
	"labix.org/v2/mgo/bson"
	"github.com/MarketReaction/sentiment-go/analyser/model"
)

func RepoFindStory(id string) Story {
	session, c, err := getMongoCollection("stories")

	defer session.Close()

	log.Printf("Finding video [%s]", id)

	result := Story{}
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		log.Printf("ERROR [%s]", err)
		panic(err)
	}

	return result
}