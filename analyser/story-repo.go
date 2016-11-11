package main

func RepoFindStory(id string) Story {
	session, c, err := getMongoCollection()

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

func getMongoCollection() (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR") + ":" + os.Getenv("MONGO_PORT_27017_TCP_PORT"))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	return session, session.DB("MarketReaction").C("stories"), err
}