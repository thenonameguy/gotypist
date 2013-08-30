package main

import(
  "log"
  "net/http"
  "code.google.com/p/go.net/websocket"
)

type Server struct{
  pattern string
  messages []*Message
  clients map[int]*Client
  addCh chan *Client
  delCh chan *Client
  sendAllCh chan *Message
  doneCh chan bool
}

func NewServer(pattern string) *Server{
  messages:=[]*Message{}
  clients:=make(map[int]*Client)
  addCh:=make(chan *Client)
  delCh:=make(chan *Client)
  sendAllCh:=make(chan *Message)
  doneCh:=make(chan bool)

  return &Server{
    pattern,
    messages,
    clients,
    addCh,
    delCh,
    sendAllCh,
    doneCh,
  }
}

func (s *Server) Add(c *Client){
  s.addCh <- c
}

func (s *Server) Del(c *Client){
  s.delCh <- c
}

func (s *Server) SendAll(msg *Message){
  s.sendAllCh <- msg
}

func (s *Server) Done(){
  s.doneCh <- true
}

func (s *Server) sendAll(msg *Message){
  for _,c:=range s.clients{
    c.Write(msg)
  }
}

func (s *Server) Listen(){
  log.Println("Server listening...")
  onConnected:=func(ws *websocket.Conn){
    defer ws.Close()
    client:=NewClient(ws,s)
    s.Add(client)
    client.Listen()
  }
  http.Handle(s.pattern,websocket.Handler(onConnected))
  for{
    select{
    case c:=<-s.addCh:
      s.clients[c.id]=c
    case c:=<-s.delCh:
      delete(s.clients,c.id)
    case msg:= <-s.sendAllCh:
      s.messages=append(s.messages,msg)
      s.sendAll(msg)
    case <-s.doneCh:
      return
    }
  }
}
