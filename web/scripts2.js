var w = window,
	d = document,
feed = d.getElementById("feed-text"),
sharing = d.getElementById("sharing"),
racercount=d.getElementById("racercount"),
type = d.getElementById("typer-input");

function newEvent(element, what, callback){
	if (element.attachEvent){
		element.attachEvent('on' + what, callback);
	}
	else {
		element.addEventListener(what, callback);
	}
}

function init_scripts(){
	var feed_text_arr = feed.innerHTML.split(" ");
	var current_word = 0;
	feed.innerHTML = feed_text_arr.slice(current_word,current_word + 6).join(" ");
	
	function check_error(e){
		if (e.keyCode == 32){
			var typeVal = type.value.split(' ')[0];
			if (typeVal == feed_text_arr[current_word]){
				if (current_word == feed_text_arr.length - 1){
					type.value = "";
					end();
					return;
				}
				++current_word;
        sendMsg({type:"word",body:current_word.toString()});
				feed.innerHTML = feed_text_arr.slice(current_word,current_word + 6).join(" ");
				type.value = "";
				type.style.background = "#fff";
			}
			else {
				type.style.background = "#e74c3c";
			}
		}
	}
	
	function end(){
		feed.innerHTML = "You win!";
	}
	newEvent(type, "keyup", check_error);
	
}

function handleMsg(msg){
  switch (msg.type)
  {
    case "text":
      feed.innerHTML=msg.body;
      init_scripts();
      break;
    case "count":
      racercount.innerHTML=msg.body;
      break;
      
  }
  console.log(msg);
}

function sendMsg(msg){
  socket.send(JSON.stringify(msg));
}
