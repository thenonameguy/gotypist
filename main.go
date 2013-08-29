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

func serveSingle(pattern string, filename string) {
  http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filename)
  })
}
func main() {
	flag.StringVar(&port, "port", ":8080", "serving port")
	flag.Parse()
	fmt.Println("Listening on port", port)
  http.HandleFunc("/index.html",IndexHandler)
	http.HandleFunc("/race", RaceHandler)
	http.Handle("/socket", websocket.Handler(SocketServer))
  serveSingle("/scripts.js","web/scripts.js")
  serveSingle("/style.css","web/style.css")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
