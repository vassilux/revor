package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func ListUsers(mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("users")
	results := []bson.M{}
	return results
}
