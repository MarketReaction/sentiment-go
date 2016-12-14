package repo

import (
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"labix.org/v2/mgo/bson"
	"log"
)

func RepoInsertStorySentiment(storySentimentToSave *model.StorySentiment) *model.StorySentiment {

	session, c, err := GetMongoCollection("storySentiment")

	defer session.Close()

	storySentimentToSave.Id = bson.NewObjectId()

	// Insert Datas
	err = c.Insert(storySentimentToSave)

	if err != nil {
		log.Printf("ERROR [%s]", err)
		panic(err)
	}

	return storySentimentToSave
}