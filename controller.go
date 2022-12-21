package main

import(
  "fmt"
)

func mainpage(w http.ResponseWriter,r *http.Request){
  http.ServeFile(w,r,"html/chat.html")
}
//reference: https://pkg.go.dev/github.com/gorilla/websocket

func sock(w http.ResponseWriter, r * http.Request){
  conn,err:=upgrader.Upgrade(w,r,nil)
  log.Fatal(err)
  for{
    msgtype,msg,err:=conn.ReadMessage()
    log.Fatal(err)
    log.Fatal(conn.WriteMessage(msgtype,msg))
  }
}
