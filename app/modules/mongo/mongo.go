package mongo

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"sync"
	"time"
)

// Extension du controlleur.
type Mongo struct {
	*revel.Controller
	MongoSession  *mgo.Session
	MongoDatabase *mgo.Database
}

type Cdr struct {
	Id             bson.ObjectId `bson:"_id"`
	Calldate       time.Time     `bson:"calldate"`
	MetadataDt     time.Time     `bson:"metadataDt"`
	ClidName       string        `bson:"clidName"`
	ClidNumber     string        `bson:"clidNumber"`
	Src            string        `bson:"src"`
	Channel        string        `bson:"channel"`
	Dcontext       string        `bson:"dcontext"`
	Disposition    int           `bson:"disposition"`
	Answerwaittime int64         `bson:"answerwaittime"`
	Billsec        int           `bson:"billsec"`
	Duration       int           `bson:"duration"`
	Uniqueid       string        `bson:"uniqueid"`
	Inoutstatus    int           `bson:"inoutstatus"`
	Recordfile     string        `bson:"recordfile"`
	Dst            string        `bson:"dst"`
}

// Stockage global de la session dont la visibilité est restreinte au package.
var session *mgo.Session

// Singleton
var dial sync.Once

// Renvoie la session mgo en cours, si aucune n'existe, elle est créée.
func GetSession() *mgo.Session {

	host, _ := revel.Config.String("mongo.host")
	// Grâce au package sync cette fonction n'est appelée
	// qu'une seule fois et de manière synchrone.
	dial.Do(func() {
		var err error
		session, err = mgo.Dial(host)
		if err != nil {
			panic(err)
		}
	})

	return session
}

// Alimente les propriétés affectées au controlleur en clonant la session mongo.
func (c *Mongo) Bind() revel.Result {
	// Oublie pas de mettre mongo.database dans le app.conf, genre "localhost"
	databaseName, _ := revel.Config.String("mongo.database")
	log.Printf("Clone the session to the mongo database : %s", databaseName)
	c.MongoSession = GetSession().Clone()
	c.MongoDatabase = c.MongoSession.DB(databaseName)

	return nil
}

// Ferme un clone
func (c *Mongo) Close() revel.Result {

	if c.MongoSession != nil {
		c.MongoSession.Close()
		log.Printf("Cloned database connexion closed.")
	}

	return nil
}

// Fonction appelée au chargement de l'application.
// Elle effectue un appel a notre fonction Bind avant
// chaque execution du controlleur.
func init() {
	revel.InterceptMethod((*Mongo).Bind, revel.BEFORE)
	revel.InterceptMethod((*Mongo).Close, revel.AFTER)
	// On veut aussi fermer le clone si le controlleur plante.
	revel.InterceptMethod((*Mongo).Close, revel.PANIC)
}
