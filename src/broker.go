package src

import (
	"net"
	"sync"
)

type publisherRegistry struct {
	publishers []string// registry for publishers
}

type subscriberRegistry struct {
	consumers []string // registry for subscribers
}


type brokerContext struct {
	pub publisherRegistry
	cons subscriberRegistry
}

type broker struct {
	l net.Listener

	sync.Mutex
	context brokerContext

	addr string // rpc address
	topics map[string][]string // topics to fileName map
	TtoS map[string][]string // topics to subscriber map
	shutdown chan struct{}
}

func NewBroker() *broker {
	return &broker{}
}

func (br *broker) Start() {

}
