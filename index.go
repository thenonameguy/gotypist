package main

import(
  "time"
  "crypto/sha1"
  "encoding/hex"
  "html/template"
  "net/http"
  "log"
)

type IndexInfo struct{
  RaceID string
}

func IndexHandler(w http.ResponseWriter,r *http.Request){
  t,err:=template.ParseFiles("web/index.html")
  if err!=nil{
    log.Fatal("Failed to parse file:",err)
  }
  info:=IndexInfo{generateRaceID()}
  t.Execute(w,info)
}

func generateRaceID() string{
  sha:=sha1.New()
  uniq:=[]byte(time.Now().String())
  sha.Write(uniq)
  return hex.EncodeToString(sha.Sum(nil))
}
