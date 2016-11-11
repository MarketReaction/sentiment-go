package repo

import (
	"log"
	"labix.org/v2/mgo/bson"
	"github.com/MarketReaction/sentiment-go/analyser/model"
)

func RepoFindSource(id string) Source {
	session, c, err := getMongoCollection("sources")

	defer session.Close()

	log.Printf("Finding video [%s]", id)

	result := Source{}
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		log.Printf("ERROR [%s]", err)
		panic(err)
	}

	return result
}