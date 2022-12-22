package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
  "github.com/gorilla/sessions"
)

var upgrader =websocket.Upgrader{
  ReadBufferSize:1024,
  WriteBufferSize:1024,
}

var store=sessions.NewCookieStore([]byte("my-name-is-niraj232"))

type Profile struct{
  uid string `json:"uid"`
  messages string `json:"messages"`
}

func cut(s string) string{
  pos:=0
  for i:=len(s)-1; i>-1;i--{
    if s[i]==':' {
      pos=i
      break
    }
  }
  s=s[pos+1:]
  return s
}

func main(){
  fs := http.FileServer(http.Dir("html/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/",mainpage)
  http.HandleFunc("/echo",sock)
  log.Println("starting server 8080")
  log.Fatal(http.ListenAndServe(":8080",nil))
}
