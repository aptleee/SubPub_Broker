package src

import (
	"encoding/json"
	"log"
	"os"
)

func (br *broker) Subscribe(topic string, subName string) Reply {
	_, ok := br.topics[topic]
	if ok == false {
		br.topics[topic] = []string{}
		br.topics[topic] = append(br.topics[topic], subName)
	} else {
		br.topics[topic] = append(br.topics[topic], subName)
	}
	return Reply{1}
}

func (br *broker)  Unsubscribe(topic string, subName string) Reply {
	_, ok := br.topics[topic]
	if ok == false {

	} else {
	}
	return Reply{1}
}

func (br *broker) Publish(topic string, msg message) Reply {
	fileName := topic
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal("can not open the file")
		return Reply{0}
	}
	enc := json.NewEncoder(f)
	err = enc.Encode(msg)
	if err != nil {
		log.Fatal("can not encode")
		return Reply{0}
	}
	defer f.Close()
	return Reply{1}
}

func (br *broker) ShutDown() {

}

func (br *broker) Pull() *MessageStream {

}

