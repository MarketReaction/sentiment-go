package main

import (
	"testing"
	"os"
	"github.com/niilo/golib/test/dockertest"
	"gopkg.in/mgo.v2"
	"log"
)

func TestAnalyser(t *testing.T) {

	// SetupMongoContainer may skip or fatal the test if docker isn't found or something goes
	// wrong when setting up the container. Thus, no error is returned
	containerID, ip := dockertest.SetupMongoContainer(t)
	defer containerID.KillRemove(t)

	os.Setenv("MONGO_PORT_27017_TCP_ADDR", ip)
	os.Setenv("MONGO_PORT_27017_TCP_PORT", "27017")

	mongoSession, err := mgo.Dial(ip)
	if err != nil {
		log.Printf("MongoDB connection failed, with address '%s'.")
	}
	defer mongoSession.Close()

	os.Args = []string{"/analyse", "12345"}

	main()
}