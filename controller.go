package main

import(
  "log"
  "net/http"
)

func mainpage(w http.ResponseWriter,r *http.Request){
//  http.ServeFile(w,r,"test.html")
  http.ServeFile(w,r,"html/chat.html")
}
//reference: https://pkg.go.dev/github.com/gorilla/websocket

func sock(w http.ResponseWriter, r * http.Request){
  conn,err:=upgrader.Upgrade(w,r,nil)
  if err != nil {
    log.Println(err)
    return
  }
  for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        log.Println(err)
        return
    }
    if err := conn.WriteMessage(messageType, p); err != nil {
        log.Println(err)
        return
    }
  } 
}
