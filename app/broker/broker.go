package broker

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"revor/app/modules/mongo"
	"time"
)

var (
	MyBroker = &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
)

type Broker struct {

	// Create a map of clients, the keys of the map are the channels
	// over which we can push messages to attached clients.  (The values
	// are just booleans and are meaningless.)
	//
	Clients map[chan string]bool

	// Channel into which new clients can be pushed
	//
	NewClients chan chan string

	// Channel into which disconnected clients should be pushed
	//
	DefunctClients chan chan string

	// Channel into which messages are pushed to be broadcast out
	// to attahed clients.
	//
	Messages chan string
}

func (b *Broker) Start() {
	// Start a goroutine
	//
	go func() {

		// Loop endlessly
		//Bytes()
		for {

			// Block until we receive from one of the
			// three following channels.
			select {

			case s := <-b.NewClients:

				// There is a new client attached and we
				// want to start sending them messages.
				b.Clients[s] = true
				revel.TRACE.Println("Added new client")

			case s := <-b.DefunctClients:

				// A client has dettached and we want to
				// stop sending them messages.
				delete(b.Clients, s)
				revel.TRACE.Println("Removed client")

			case msg := <-b.Messages:

				// There is a new message to send.  For each
				// attached client, push the new message
				// into the client's message channel.
				for s, _ := range b.Clients {
					s <- msg
				}
				revel.TRACE.Printf("Broadcast message to %d clients", len(b.Clients))
			}
		}
	}()
	revel.INFO.Println(" [Broker] Started")

}

//Broadcaster of new data for realtime notification
//Redis queue can be used as realtime source notificateur
func processFetchEvens() {
	session, db := mongo.GetDatabase()
	timenow := time.Now()
	var currentDateString = timenow.Format(time.RFC3339)
	didCalls := mongo.GetDidCalls(currentDateString, db)
	peerInCalls := mongo.GetPeerInCalls(currentDateString, db)
	peerOutCalls := mongo.GetPeerOutCalls(currentDateString, db)
	results := bson.M{
		"didCalls":     didCalls,
		"peerInCalls":  peerInCalls,
		"peerOutCalls": peerOutCalls,
	}
	data, err := json.Marshal(results)
	if err == nil {
		MyBroker.Messages <- fmt.Sprintf("%s", data)
	}
	if session != nil {
		session.Close()
	}
	revel.TRACE.Printf("Broadcast message [%s] to %d clients", data, len(MyBroker.Clients))
}

func broker() {
	revel.INFO.Println(" [Broker] Initialize broker")
	MyBroker.Start()
	go func() {
		for i := 0; ; i++ {
			// Create a message to send to clients,
			// including the current time.
			if len(MyBroker.Clients) > 0 {
				processFetchEvens()
			}
			time.Sleep(2 * 1e9)

		}
	}()
}

func init() {
	go broker()
}
