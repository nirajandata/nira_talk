let socket = new WebSocket("ws://localhost:8080/echo");
let url="http://localhost:8080/encode" 

var view=document.getElementById("msglist")
var sent=document.getElementById("msgbox")
var btn=document.getElementById("btn")

var you="",uid="",messages,pmsg="";

async function jsonloader(){
  const response=await fetch(url);
  const result=await response.json();
  uid=result.uid,messages=result.messages;
  if(uid!=you && messages!=pmsg){
  prints(uid,messages)
  pmsg=messages;
  }
}

setInterval(jsonloader,1000);
console.log(uid,messages);
socket.onopen=()=>{
  socket.send("joins the chat")
};
socket.onmessage=(e)=>{
  if(you=="" && uid!="")
    you=uid;
  prints(uid,messages);
};
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
function prints(uid,msg){
  view.innerText+=uid+" "+msg +'\n';
}
socket.onclose=(e)=>{
  socket.send("thanks for using it ")
}
