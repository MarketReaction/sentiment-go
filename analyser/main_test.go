package main

import (
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"github.com/MarketReaction/sentiment-go/analyser/repo"
	"gopkg.in/ory-am/dockertest.v2"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestAnalyser_withNoNamedEntities_Stops(t *testing.T) {

	var id bson.ObjectId = bson.NewObjectId()

	story := &model.Story{
		Id:    id,
		Title: "Test Story",
	}

	var db *mgo.Session
	var c dockertest.ContainerID
	var apic dockertest.ContainerID

	db, apic, c = setUp()
	addTestStory(story)

	defer db.Close()
	defer apic.KillRemove()
	defer c.KillRemove()

	os.Args = []string{"/analyse", id.Hex()}

	main()
}

func TestAnalyser_withNamedEntities_CallsSentimentApi(t *testing.T) {

	var storyId bson.ObjectId = bson.NewObjectId()
	var matchedCompanyId bson.ObjectId = bson.NewObjectId()

	story := &model.Story{
		Id:    storyId,
		Title: "Test Story",
		NamedEntities: model.NamedEntities{
			Organisations: []model.NamedEntity{
				{
					Name:    "Test Name",
					Matched: true,
					Sentiments: []model.Sentiment{
						{
							Sentence: "This is a really good sentence",
						},
					},
				},
			},
		},
		MatchedCompanies: []string{matchedCompanyId.Hex()},
	}

	company := &model.Company{
		Id:    matchedCompanyId,
		Name:	"Test Company",
	}

	var db *mgo.Session
	var c dockertest.ContainerID
	var apic dockertest.ContainerID

	db, apic, c = setUp()
	addTestStory(story)
	addTestCompany(company)

	defer db.Close()
	defer apic.KillRemove()
	defer c.KillRemove()

	os.Args = []string{"/analyse", storyId.Hex()}

	main()
}

func setUp() (*mgo.Session, dockertest.ContainerID, dockertest.ContainerID) {

	apic, ip, port, err := dockertest.SetupCustomContainer("marketreaction/sentiment-api", 8888, 10*time.Second)
	if err != nil {
		log.Fatalf("Could not setup container: %s", err)
	}

	log.Printf("Sentiment API started at http://%s:%d", ip, port)

	const delay = 5 * time.Second
	time.Sleep(delay)

	os.Setenv("SENTIMENT_API_ADDR", ip)
	os.Setenv("SENTIMENT_API_PORT", strconv.Itoa(port))

	var db *mgo.Session
	c, err := dockertest.ConnectToMongoDB(15, time.Millisecond*500, func(url string) bool {
		// This callback function checks if the image's process is responsive.
		// Sometimes, docker images are booted but the process (in this case MongoDB) is still doing maintenance
		// before being fully responsive which might cause issues like "TCP Connection reset by peer".
		var err error
		db, err = mgo.Dial(url)
		if err != nil {
			return false
		}

		// Sometimes, dialing the database is not enough because the port is already open but the process is not responsive.
		// Most database conenctors implement a ping function which can be used to test if the process is responsive.
		// Alternatively, you could execute a query to see if an error occurs or not.
		return db.Ping() == nil
	})

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	var mongoHost []string = strings.Split(db.LiveServers()[0], ":")

	os.Setenv("MONGO_PORT_27017_TCP_ADDR", "127.0.0.1")
	os.Setenv("MONGO_PORT_27017_TCP_PORT", mongoHost[1])

	return db, apic, c
}

func addTestStory(story *model.Story) {
	session, con, err := repo.GetMongoCollection("stories")

	if err != nil {
		log.Fatalf("Failed to insert test story: %s", story.Id)
	}

	defer session.Close()

	log.Println("Inserting story for Test")

	con.Insert(story)

	return
}

func addTestCompany(company *model.Company) {
	session, con, err := repo.GetMongoCollection("companies")

	if err != nil {
		log.Fatalf("Failed to insert test company: %s", company.Id)
	}

	defer session.Close()

	log.Println("Inserting story for Test")

	con.Insert(company)

	return
}