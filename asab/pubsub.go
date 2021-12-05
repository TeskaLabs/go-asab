package asab

import "sync"

// Inspired by https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/

type PubSub struct {
	_Mutex       sync.RWMutex
	_Subscribers map[string][]func(PubSubMessage)
}

type PubSubMessage struct {
	Name string
	I    interface{}
}

func (ps *PubSub) Initialize() {
	ps._Subscribers = make(map[string][]func(PubSubMessage))
}

func (ps *PubSub) Subscribe(message_type string, handler func(PubSubMessage)) {
	ps._Mutex.Lock()
	defer ps._Mutex.Unlock()

	ps._Subscribers[message_type] = append(ps._Subscribers[message_type], handler)
}

//TODO: func (ps *PubSub) Unsubscribe(message_type string, callback func())

func (ps *PubSub) Publish(message PubSubMessage) {
	ps._Mutex.Lock()
	defer ps._Mutex.Unlock()

	for _, handler := range ps._Subscribers[message.Name] {
		handler(message)
	}
}
