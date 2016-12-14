package main

import (
	"encoding/json"
	"fmt"
	"github.com/MarketReaction/sentiment-go/analyser/model"
	"log"
	"net/http"
	"net/url"
	"os"
)

type SentimentApiResponse struct {
	Score int `json:"score"`
}

func Analyse(namedEntities model.NamedEntities) model.NamedEntities {

	log.Println("Analysing Named Entities")
	log.Output(0, namedEntities)

	for i, org := range namedEntities.Organisations {
		if org.Matched {
			for is, sent := range org.Sentiments {
				namedEntities.Organisations[i].Sentiments[is].Sentiment = getScoreForText(sent.Sentence)
			}
		}
	}

	log.Output(1, namedEntities)

	return namedEntities
}

func getScoreForText(text string) int {
	apiUrl := fmt.Sprintf("http://%s:%s", os.Getenv("SENTIMENT_API_ADDR"), os.Getenv("SENTIMENT_API_PORT"))

	form := url.Values{}
	form.Add("text", text)

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.PostForm(apiUrl, form)
	if err != nil {
		log.Fatal("Do Failed: ", err)
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record SentimentApiResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	log.Printf("Score: %d", record.Score)

	return record.Score
}
