package repo

import (
	"os"
	"labix.org/v2/mgo"
)

func getMongoCollection(collection string) (*mgo.Session, *mgo.Collection, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR") + ":" + os.Getenv("MONGO_PORT_27017_TCP_PORT"))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	return session, session.DB("MarketReaction").C(collection), err
}