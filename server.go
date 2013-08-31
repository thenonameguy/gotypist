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
      log.Println(c.id,"disconnected from race id:",r.id,". Remaining players:",len(r.clients)-1)
			delete(r.clients, c.id)
      if(len(r.clients)==0){
        log.Println("No players left in race",r.id,"! Deleting...")
        delete(r.server.races,r.id)
        return
      } 
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
    if _,created:=s.races[raceid];!created{
      s.races[raceid]=NewRace(s)
      log.Println("Created race: ",raceid)
      s.races[raceid].id=raceid
      go s.races[raceid].RaceSocketHandler()
    }
		client := NewClient(ws, s.races[raceid])
    s.races[raceid].Add(client)
    log.Println("Client",client.id,"connected to race",raceid)
    client.Write(&Message{"text",s.races[raceid].txt})
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
}
