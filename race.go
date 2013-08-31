package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type Race struct{
  txt string
  id string
  started bool
  server    *Server
	clients   map[uint64]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
}

func NewRace(s *Server) *Race{
	clients := make(map[uint64]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
  return &Race{chooseText(),"",false,s,clients,addCh,delCh,sendAllCh}
}

func (r *Race) Add(c *Client) {
	r.addCh <- c
}

func (r *Race) Del(c *Client) {
	r.delCh <- c
}

func (r *Race) SendAll(msg *Message) {
	r.sendAllCh <- msg
}

func (r *Race) sendAll(msg *Message) {
	for _, c := range r.clients {
		c.Write(msg)
	}
}

func chooseText() string {
	files, _ := ioutil.ReadDir("txts")
	content, err := ioutil.ReadFile("txts/" + files[rand.Int31n(int32(len(files)))].Name())
	if err != nil {
		log.Println("txt error", err)
	}
	return string(content)
}

func getRaceID(r *http.Request) (string, error) {
	if raceid := r.FormValue("id"); raceid == "" {
		return "", errors.New("race: no id got from GET request")
	} else {
		return raceid, nil
	}
}
func RaceHandler(w http.ResponseWriter, r *http.Request) {
	raceid, err := getRaceID(r)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Race 404'd, go back"))
		return
	}
	t, err := template.ParseFiles("web/race.html")
	if err != nil {
		log.Fatal("Failed to parse file:", err)
	}
	t.Execute(w,raceid)
}
