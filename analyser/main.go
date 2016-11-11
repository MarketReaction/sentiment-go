package main // import "github.com/MarketReaction/sentiment-go/analyser"

import (
	"os"
	"fmt"
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"github.com/MarketReaction/sentiment-go/analyser/repo"
)

func main() {

    	args := os.Args[1:]

	fmt.Println("Args: " + args)

	var storyId string = args[0]

	fmt.Println("StoryId: " + storyId)

	story := RepoFindStory(storyId)

	fmt.Println("StoryTitle: " + story.title)
}
