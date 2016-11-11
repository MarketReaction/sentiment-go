package model

type NamedEntities struct {
	organisations NamedEntity       `bson:"organisations"`
	people        NamedEntity       `bson:"people"`
	misc          NamedEntity       `bson:"misc"`
	locations     NamedEntity       `bson:"locations"`
}
