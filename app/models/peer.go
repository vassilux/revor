package models

import (
	"fmt"
)

type Peer struct {
	Peer    string `bson:"peer"			   json:"peer"`
	Value   string `json:"value"			bson:"value"`
	Comment string `json:"comment"			bson:"comment"`
}

func (p *Peer) String() string {
	return fmt.Sprintf("Peer(%s)", p.Peer)
}
