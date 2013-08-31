package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
  "strconv"
)

const bufferSize = 100

var maxId uint64 = 0

type Client struct {
	id   uint64
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
		log.Println("Client", c.id, "disconnected")
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
			log.Println("Sending to", c.id, ":", msg)
		}
	}
}

func (c *Client) listenRead() {
	for {
		select {
		default:
			var msg Message
			if err := websocket.JSON.Receive(c.ws, &msg);err!=nil{
        if err.Error()=="EOF"{
          c.race.Del(c)
          return
        }
        log.Println("Unhandled error:",err)
        return
      }
			log.Println(c.id, "got:", msg)

      if msg.Type=="word"{
        i,err:=strconv.Atoi(msg.Body)
        if err!=nil{
          log.Println(err)
          continue
        }
        if i>c.race.best{
          c.race.best=i
          c.race.sendAll(&Message{"best",msg.Body})
        }
      }
		}
	}
}
