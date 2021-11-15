package websocket

import (
	"errors"
	"sync"
)

type Pool interface {
	// Register adds the client to be broadcasted to using an id and channel from which to receive messages to broadcast
	Register(clientID string, receiver chan<- Message) error
	// Unregister removes the client from the pool and will no longer receive messages to broadcast
	Unregister(id string)
	// Broadcast tells the client to pass a message to all other clients to send a message
	Broadcast(clientID string, msg Message)
}

type pool struct {
	clients map[string]chan<- Message
	lock    *sync.RWMutex
	exit    chan<- interface{}
}

var (
	ErrorClientIDAlreadyRegistered = errors.New("ClientID is already registered with the pool")
)

const (
	ExitMessageType = "Pool Exiting"
	exitMessage     = Message{Type: ExitMessageType}
)

func NewPool() (*pool, <-chan interface{}) {
	exit := make(chan interface{})
	p := &pool{
		clients: make(map[string]chan<- Message),
		lock:    &sync.RWMutex{},
		exit:    exit,
	}

	p.Start()

	return p, exit
}

func (p *pool) Register(clientID string, receiver chan<- Message) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if _, ok := p.clients[clientID]; ok {
		return ErrorClientIDAlreadyRegistered
	}

	p.clients[clientID] = receiver

	return nil
}

func (p *pool) UnRegister(clientID string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.clients, clientID)
}

func (p *pool) Broadcast(clientID string, msg Message) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	for id, receiver := range p.clients {
		if id != clientID {
			receiver <- msg
		}
	}
}

func (p *pool) Start() {
	for {
		select {
		case <-p.exit:
			p.lock.Lock()
			for id, receiver := range p.clients {
				receiver <- exitMessage
				delete(p.clients, id)
				close(receiver)
			}
			p.lock.Unlock()
			break
		}
	}
}
