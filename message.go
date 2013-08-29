package main

type Message struct{
  Type string `json:"type"`
  Body string `json:"body"`
}

func (m *Message) String() string{
  return m.Type+": "+m.Body
}
