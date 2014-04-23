package mongo

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"log"
	"sync"
)

const (
	DB_REQUEST_LIMITS = 10000
)

// Controller extended.
type Mongo struct {
	*revel.Controller
	MongoSession  *mgo.Session
	MongoDatabase *mgo.Database
}

// Global session for package.
var session *mgo.Session

// Singleton
var dial sync.Once

// Return the mongo current session, if the session not existe create one new.
func GetSession() *mgo.Session {
	host, _ := revel.Config.String("mongo.host")
	dial.Do(func() {
		var err error
		session, err = mgo.Dial(host)
		if err != nil {
			panic(err)
		}
	})

	return session
}

//
func (c *Mongo) Bind() revel.Result {
	//
	databaseName, _ := revel.Config.String("mongo.database")
	log.Printf("Clone the session to the mongo database : %s", databaseName)
	c.MongoSession = GetSession().Clone()
	c.MongoDatabase = c.MongoSession.DB(databaseName)

	return nil
}

//
func (c *Mongo) Close() revel.Result {

	if c.MongoSession != nil {
		c.MongoSession.Close()
		log.Printf("Cloned database connexion closed.")
	}

	return nil
}

func GetDatabase() (*mgo.Session, *mgo.Database) {
	databaseName, _ := revel.Config.String("mongo.database")
	log.Printf("GetDatabase Clone the session to the mongo database : %s", databaseName)
	session := GetSession().Clone()
	db := session.DB(databaseName)
	return session, db
}

// Module initialization
func init() {
	revel.InterceptMethod((*Mongo).Bind, revel.BEFORE)
	revel.InterceptMethod((*Mongo).Close, revel.AFTER)
	// oups close session of the contoler get a panic
	revel.InterceptMethod((*Mongo).Close, revel.PANIC)
}
