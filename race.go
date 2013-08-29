package main

import(
  "html/template"
  "io/ioutil"
  "net/http"
  "errors"
  "log"
)

type RaceInfo struct{
  Txt string
  Racers map[int]*Racer
}

type Racer struct{
  ID int
}

var races = make(map[string]*RaceInfo)

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

func getRaceID(r *http.Request) (string,error){
  if raceid:=r.FormValue("id");raceid=="" {
    return "",errors.New("race: no id got from GET request")
  } else{
    return raceid,nil
  }
}

func Join(id string) *Racer{
  //TODO: Inform all other racers that someone joined
  r:=new(Racer)
  r.ID=len(races[id].Racers)
  races[id].Racers[r.ID]=r
  return r
}

func RaceHandler(w http.ResponseWriter,r *http.Request){
  raceid,err:=getRaceID(r)
  if err!=nil{
    log.Println(err)
    w.Write([]byte("Race 404'd, go back"))
    return
  }
  if len(races[raceid].Racers)==0{
    races[raceid].Racers=make(map[int]*Racer)
  }
  t,err:=template.ParseFiles("web/race.html")
  if err!=nil{
    log.Fatal("Failed to parse file:",err)
  }
  racer:=Join(raceid)
  log.Println(len(races[raceid].Racers),racer)
  t.Execute(w,races[raceid])
}
