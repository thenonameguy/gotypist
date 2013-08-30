package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port string

func serveFile(pattern string, filename string) {
  http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, filename)
  })
}

func main() {
	flag.StringVar(&port, "port", ":8080", "serving port")
	flag.Parse()
	fmt.Println("Listening on port", port)
  http.HandleFunc("/",IndexHandler)
	http.HandleFunc("/race", RaceHandler)
  serveFile("/scripts.js","web/scripts.js")
  serveFile("/scripts2.js","web/scripts2.js")
  serveFile("/style.css","web/style.css")
  serveFile("/bg.png","web/bg.png")
  server:=NewServer("/socket")
  go server.Listen()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
