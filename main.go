package main

import (
//	"encoding/json"
	"flag"
	"fmt"
  "code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func SocketServer(ws *websocket.Conn){

}

var port string

func main() {
	flag.StringVar(&port, "port", ":8080", "serving port")
	flag.Parse()
	fmt.Println("Listening on port", port)
  http.HandleFunc("/index.html",IndexHandler)
	http.HandleFunc("/race", RaceHandler)
	http.Handle("/socket", websocket.Handler(SocketServer))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
