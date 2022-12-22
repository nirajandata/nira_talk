package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
  "github.com/gorilla/sessions"
)

var upgrader =websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true // Accepting all requests
  },

  ReadBufferSize:1024,
  WriteBufferSize:1024,
}

var store=sessions.NewCookieStore([]byte("test"))

type Profile struct{
  Uid string `json:"uid"`
  Messages string `json:"messages"`
}

var api Profile

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
  http.HandleFunc("/encode",jsonthrow)
  log.Println("starting server 8080")
  log.Fatal(http.ListenAndServe(":8080",nil))
}
