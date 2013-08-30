package main

import(
  "html/template"
  "io/ioutil"
  "net/http"
  "errors"
  "log"
)

func ChooseText() string{
  files,err:=ioutil.ReadDir("web")
  if err!=nil{
    log.Println("Couldn't access directory:",err)
    return ""
  }
  for _,v:=range files{
    log.Println(v)
  }
  return "lel"
}

func RaceHandler(w http.ResponseWriter,r *http.Request){
  raceid,err:=getRaceID(r)
  if err!=nil{
    log.Println(err)
    w.Write([]byte("Race 404'd, go back"))
    return
  }
  t,err:=template.ParseFiles("web/race.html")
  if err!=nil{
    log.Fatal("Failed to parse file:",err)
  }
  racer:=Join(raceid)
  log.Println(len(races[raceid].Racers),racer)
  t.Execute(w,ChooseText())
}
