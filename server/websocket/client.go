package websocket

import (
	"errors"
	"fmt"

	"github.com/driabwb/retroboard/application"
	"github.com/driabwb/retroboard/messages"
	"github.com/google/uuid"
)

type Client interface {
	Send(messages.Message) error
}

type WebsocketConn interface {
	Send(messages.Message) error
	NextMessage() (messages.Message, error)
	Close() error
}

type client struct {
	id                string
	conn              WebsocketConn
	readReceiver      <-chan messages.Message
	broadcastReceiver <-chan messages.Message
	pool              Pool
	app               *application.App
}

func NewClient(conn WebsocketConn, pool Pool, app *application.App) (*client, error) {
	id := uuid.New()
	readChan := make(chan messages.Message)
	broadcastChan := make(chan messages.Message)

	err := pool.Register(id, broadcastChan)

	for err != nil {
		if !errors.Is(err, ErrorClientIDAlreadyRegistered) {
			return nil, fmt.Errorf("Failed to register client with the pool: %w", err)
		}

		id = uuid.New()
		err = pool.Register(id, broadcastChan)
	}

	newClient := &client{
		id:                id,
		conn:              conn,
		readReceiver:      readChan,
		broadcastReceiver: broadcastChan,
		pool:              pool,
		app:               app,
	}

	go messageReceiver(conn, readChan)

	go newClient.start()

	return newClient, nil
}

func (c *client) start() {
	for {
		select {
		case msg := <-broadcastReceiver:
			if msg.Type == ExitMessageType {
				// send close message

				// Do NOT unregister because the pool is closing
				break
			}
			// Pass along the message
		case msg := <-readReceiver:
			isClosing := c.handleMessage(msg)
			if isClosing {
				// Unregister as the client is closing
				c.pool.Unregister(c.id)
				break
			}
		}
	}

	// When exiting close connection
	// TODO: handle error
	conn.Close()
}

func (c *client) Send(msg messages.Message) error {
	// TODO: Coordinate sending messages?
	return c.conn.Send(msg)
}

func messageReceiver(conn WebsocketConn, msgChan chan<- messages.Message) {
	for {
		// Will return an error when conn is closed
		msg, err := conn.Read()

		// if err send error close loop
		if err != nil {
			// TODO: log errors and maybe be more technical and handle issues
			break
		}

		msgChan <- msg
	}

	close(msgChan)
}
