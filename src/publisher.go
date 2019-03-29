package src

import "sync"

type publisher struct {
	name string // RPC address, IP address
	sync.Mutex
}

func NewPublisher() *publisher {
	return &publisher{

	}
}

func (pub *publisher) Publish(topic string, msg message) Reply {
	return Reply{1}
}