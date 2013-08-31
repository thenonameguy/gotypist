package main

import "strconv"

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
  ID   uint64 `json:"id"`
}

func (m *Message) String() string {
  return "Type: "+m.Type + ", Body: " + m.Body + ", ID: "+strconv.FormatUint(m.ID,10)
}
