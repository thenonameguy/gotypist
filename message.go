package main

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

func (m *Message) String() string {
  return "Type: "+m.Type + ", Body: " + m.Body
}
