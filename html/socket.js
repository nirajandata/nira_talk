let socket = new WebSocket("ws://localhost:8080/echo");

var view=document.getElementById("msglist")
var sent=document.getElementById("msgbox")
var btn=document.getElementById("btn")


socket.onopen=()=>{
  socket.send("hello there !")
};

socket.onmessage=(e)=>{
  view.innerText +="user xyz sent " + e.data+"\n";
};


sent.addEventListener("keydown",(event)=>{
  console.log(event.shiftKey)
  if(event.code=="Enter" && !event.shiftKey) {
    chore();
  }
});
function chore(){
  socket.send(sent.value)
  sent.value='';
}

socket.onclose=(e)=>{
  socket.send("thanks for using it ")
}
