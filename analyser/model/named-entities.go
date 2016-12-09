package model

import "reflect"

type NamedEntities struct {
	Organisations []NamedEntity `bson:"organisations"`
	People        []NamedEntity `bson:"people"`
	Misc          []NamedEntity `bson:"misc"`
	Locations     []NamedEntity `bson:"locations"`
}

func (s NamedEntities) IsEmpty() bool {
	return reflect.DeepEqual(s, NamedEntities{})
}
