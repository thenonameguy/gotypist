package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

const bufferSize = 100

var maxId int = 0

type Client struct {
	id   int
	ws   *websocket.Conn
	race *Race
	ch   chan *Message
}

func NewClient(ws *websocket.Conn, race *Race) *Client {
	maxId++
	ch := make(chan *Message, bufferSize)
	return &Client{maxId, ws, race, ch}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *Message) {
	select {
	case c.ch <- msg:
	default:
		c.race.Del(c)
		log.Println("client", c.id, "disconnected")
	}
}

func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	for {
		select {
		case msg := <-c.ch:
			websocket.JSON.Send(c.ws, msg)
			log.Println("sending to", c.id, ":", msg)
		}
	}
}

func (c *Client) listenRead() {
	for {
		select {
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err != nil {
				log.Println(c.id, "error", err)
				c.race.Del(c)
        return
			}
			log.Println(c.id, "got:", msg)
			c.race.SendAll(&msg)
		}
	}
}
