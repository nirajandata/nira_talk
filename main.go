package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader =websocket.Upgrader{
  ReadBufferSize:1024,
  WriteBufferSize:1024,
}

func main(){
  fs := http.FileServer(http.Dir("html/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/",mainpage)
  http.HandleFunc("/echo",sock)
  log.Println("starting server 8080")
  log.Fatal(http.ListenAndServe(":8080",nil))
}
