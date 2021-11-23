package asab

import "sync"

// Inspired by https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/

type PubSub struct {
	_Mutex       sync.RWMutex
	_Subscribers map[string][]func()
}

func (ps *PubSub) Initialize() {
	ps._Subscribers = make(map[string][]func())
}

func (ps *PubSub) Subscribe(message_type string, callback func()) {
	ps._Mutex.Lock()
	defer ps._Mutex.Unlock()

	ps._Subscribers[message_type] = append(ps._Subscribers[message_type], callback)
}

//TODO: func (ps *PubSub) Unsubscribe(message_type string, callback func())

func (ps *PubSub) Publish(message_type string) {
	ps._Mutex.Lock()
	defer ps._Mutex.Unlock()

	for _, callback := range ps._Subscribers[message_type] {
		callback()
	}
}
