package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

type Server struct {
	pattern   string
  races     map[string]*Race
}

func NewServer(pattern string) *Server {
	races := make(map[string]*Race)

	return &Server{
		pattern,
    races,
	}
}

func (r *Race) RaceSocketHandler(){
	for {
		select {
		case c := <-r.addCh:
			r.clients[c.id] = c
		case c := <-r.delCh:
			delete(r.clients, c.id)
		case msg := <-r.sendAllCh:
			r.sendAll(msg)
		}
	}
}

func (s *Server) Listen() {
	log.Println("Server listening...")
	onConnected := func(ws *websocket.Conn) {
		defer ws.Close()
    raceid,err:=getRaceID(ws.Request())
    if err!=nil{
      websocket.JSON.Send(ws,&Message{"error","wrong url, no GET iD"})
    }
    if _,ok:=s.races[raceid];!ok{
      s.races[raceid]=NewRace()
      log.Println("Created race: ",raceid)
      go s.races[raceid].RaceSocketHandler()
    }
		client := NewClient(ws, s.races[raceid])
    s.races[raceid].Add(client)
    log.Println("Client",client.id,"connected to race",raceid)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
}
