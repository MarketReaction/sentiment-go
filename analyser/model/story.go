package model

import (
	"time"
	"labix.org/v2/mgo/bson"
)

type Story struct {
	Id               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	url              string        `json:"url"`
	dateFound        time.Time     `json:"dateFound"`
	Title            string        `json:"title"`
	body             string        `json:"body"`
	datePublished    time.Time     `json:"datePublished"`
	namedEntities 	 NamedEntities  `bson:"keywords"`

	parentSource     string        `json:"parentSource"`
	matchedCompanies []string      `json:"matchedCompanies"`
	sentiment        int           `json:"sentiment"`
}

func (f Story) Url() string {
	return f.url
}

func (f Story) ParentSource() string {
	return f.parentSource
}
func (f Story) NamedEntities() NamedEntities {
	return f.namedEntities
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