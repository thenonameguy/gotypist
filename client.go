package main

import(
  "log"
  "code.google.com/p/go.net/websocket"
)

const bufferSize=100
var maxId int = 0

type Client struct{
  id int
  ws *websocket.Conn
  server *Server
  ch chan *Message
  doneCh chan bool
}

func NewClient(ws *websocket.Conn, server *Server) *Client{
  maxId++
  ch:=make(chan *Message,bufferSize)
  doneCh:=make(chan bool)
  return &Client{maxId,ws,server,ch,doneCh}
}

func (c *Client) Conn() *websocket.Conn{
  return c.ws
}

func (c *Client) Write(msg *Message){
  select{
  case c.ch <- msg:
  default:
    c.server.Del(c)
    log.Println("client",c.id,"disconnected")
  }
}


func (c *Client) Done(){
  c.doneCh <- true
}

func (c *Client) Listen(){
  go c.listenWrite()
  c.listenRead()
}

func (c *Client) listenWrite(){
  for{
    select{
    case msg:=<-c.ch:
      websocket.JSON.Send(c.ws,msg)
    }
  }
}

func (c *Client) listenRead(){
  for{
    select{
    default:
      var msg Message
      err:=websocket.JSON.Receive(c.ws,&msg)
      if err!=nil{
        log.Println("error",err)
      }
      c.server.SendAll(&msg)
    }
  }
}
