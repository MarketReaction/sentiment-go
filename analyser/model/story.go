package model

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Story struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Url           string        `json:"url"`
	DateFound     time.Time     `json:"dateFound"`
	Title         string        `json:"title"`
	Body          string        `json:"body"`
	DatePublished time.Time     `json:"datePublished"`
	NamedEntities NamedEntities `bson:"keywords"`

	ParentSource     string   `json:"parentSource"`
	MatchedCompanies []string `json:"matchedCompanies"`
	Sentiment        int      `json:"sentiment"`
}

type Stories []Story

//    private URL url;
//    private Date dateFound;
//    private String title;
//    private String body;
//    private Date datePublished;
//    private NamedEntities entities;
//
//    private String parentSource;
//
//    private List<String> matchedCompanies;
//
//    private int sentiment;
//
//    private List<Metric> metrics;
