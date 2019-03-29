package SubPub_Broker


// message format
type message struct {
	content string
}


type Reply struct {
	r int
}

type MessageStream struct {

}

func (ms *MessageStream) iterator() *Iterator {
	return &Iterator{}
}

type Iterator struct {

}

func (iter *Iterator) Next() {

}

func (iter *Iterator) HasNext() {

}
