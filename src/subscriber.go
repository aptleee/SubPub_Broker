package src

import "sync"

type subscriber struct {
	name string
	sync.Mutex
	topics []string // which topics the sub subscribed
}

func NewSubscriber() *subscriber {
	return &subscriber{}
}


func (sub *subscriber) Subscribe(topic string) Reply {
	return Reply{1}
}

func (sub *subscriber) Pull() *MessageStream {
	return &MessageStream{}
}