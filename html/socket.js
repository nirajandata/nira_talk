let socket = new WebSocket("ws://localhost:8080/echo");

var view=document.getElementById("msglist")
var sent=document.getElementById("msgbox")
var btn=document.getElementById("btn")


socket.onopen=()=>{
  socket.send(" namaste !")
};
let user=" you"
socket.onmessage=(e)=>{
  view.innerText +=user+ " "+ e.data+"\n";
};
let uid="";
sent.addEventListener("keydown",(event)=>{
  if(event.code=="Enter" && !event.shiftKey) {
    chore();
  }
  let count=String(sent.value).length;
  if(count>90 ){
    sent.style.backgroundColor="yellow";
    sent.style.borderColor="red";
  }
  else {
    sent.style.backgroundColor="lightgreen";
  }
});
function chore(){
  socket.send(sent.value)
  uid=document.cookie;
  sent.value='';
}

socket.onclose=(e)=>{
  socket.send("thanks for using it ")
}
