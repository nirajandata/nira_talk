package main

import (
  "log"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader =websocket.Upgrader{
  ReadBufferSize:1024
  WriteBufferSize:1024
}

func main(){
  mux:=http.NewServeMux()
  mux.HandleFunc("/",mainpage)
  mux.HandleFunc("/sock",connect)
  log.Println("starting server 8080")
  log.Fatal(http.ListenAndServe(":8080",mux))
}
