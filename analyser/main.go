package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"os"
	"fmt"
	"gopkg.in/mgo.v2"
)

func main() {

    	args := os.Args[1:]

	fmt.Println(args)

	session, err := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR") + ":" + os.Getenv("MONGO_PORT_27017_TCP_PORT"))
        if err != nil {
                panic(err)
        }
        defer session.Close()

}
