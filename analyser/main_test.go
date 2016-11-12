package main

import (
	"testing"
	"os"
	"gopkg.in/ory-am/dockertest.v2"
	"gopkg.in/mgo.v2"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
	"strings"
)

func TestAnalyser(t *testing.T) {

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

	// Close db connection and kill the container when we leave this function body.
	defer db.Close()
	defer c.KillRemove()

	var mongoHost []string = strings.Split(db.LiveServers()[0], ":")

	os.Setenv("MONGO_PORT_27017_TCP_ADDR", "127.0.0.1")
	os.Setenv("MONGO_PORT_27017_TCP_PORT", mongoHost[1])

	var id string = bson.NewObjectId().Hex()

	os.Args = []string{"/analyse", id}

	main()
}