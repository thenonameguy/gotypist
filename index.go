package main

import(
  "time"
  "crypto/sha1"
  "encoding/hex"
  "html/template"
  "net/http"
  "log"
)

func IndexHandler(w http.ResponseWriter,r *http.Request){
  t,err:=template.ParseFiles("web/index.html")
  if err!=nil{
    log.Fatal("Failed to parse file:",err)
  }
  t.Execute(w,generateRaceID())
}

func generateRaceID() string{
  sha:=sha1.New()
  uniq:=[]byte(time.Now().String())
  sha.Write(uniq)
  return hex.EncodeToString(sha.Sum(nil))
}
