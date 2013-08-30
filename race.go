package main

import(
  "html/template"
  "net/http"
  "errors"
  "log"
  "io/ioutil"
  "math/rand"
)

func chooseText() string{
  files,_:=ioutil.ReadDir("txts")
  content,err:=ioutil.ReadFile("txts/"+files[rand.Int31n(int32(len(files)))].Name())
  if err!=nil{
    log.Println("txt error",err)
  }
  return string(content)
}

func getRaceID(r *http.Request) (string,error){
  if raceid:=r.FormValue("id");raceid=="" {
    return "",errors.New("race: no id got from GET request")
  } else{
    return raceid,nil
  }
}
func RaceHandler(w http.ResponseWriter,r *http.Request){
  raceid,err:=getRaceID(r)
  _=raceid
  if err!=nil{
    log.Println(err)
    w.Write([]byte("Race 404'd, go back"))
    return
  }
  t,err:=template.ParseFiles("web/race.html")
  if err!=nil{
    log.Fatal("Failed to parse file:",err)
  }
  t.Execute(w,chooseText())
}
