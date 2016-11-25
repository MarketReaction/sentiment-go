package repo

import (
	"log"
	"labix.org/v2/mgo/bson"
	"../model"
)

func RepoFindCompany(id string) *model.Company {
	session, c, err := GetMongoCollection("companies")

	defer session.Close()

	log.Printf("Finding Company [%s]", id)

	result := &model.Company{}
	err = c.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		log.Printf("ERROR [%s]", err)
		panic(err)
	}

	return result
}