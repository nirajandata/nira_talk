package main

import(
  "log"
  "net/http"
  "encoding/json"
  "strconv"
)

func mainpage(w http.ResponseWriter,r *http.Request){
//  http.ServeFile(w,r,"test.html")
  http.ServeFile(w,r,"html/chat.html")
}
//reference: https://pkg.go.dev/github.com/gorilla/websocket

func sock(w http.ResponseWriter, r * http.Request){
  conn,err:=upgrader.Upgrade(w,r,nil)
  session,_:=store.Get(r,"cookie-name")
  if err != nil {
    log.Println(err)
    return
  }
  for {
    messageType, value, err := conn.ReadMessage()
    if err != nil {
        log.Println(err)
//        return
    }
    session.Values["userid"]=cut(conn.RemoteAddr().String()) 
    session.Save(r,w)
    log.Println(session.Values["userid"])
    //todo json api
    location:=cut(conn.RemoteAddr().String())
    _,ok:=counter[location]
    if(!ok){
      count++
      counter[location]="1";
    }  
    location=strconv.Itoa(count)
    log.Println(location)
    api=Profile{
      Uid:location,
      Messages:string(value),
    } 
    /// i need to add cookie of only userid
    if err := conn.WriteMessage(messageType, value); err != nil {
        log.Println(err)
        return
    }
  } 
}
func jsonthrow(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(api)
}
