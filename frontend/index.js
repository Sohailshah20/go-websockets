const selectedChat = "general";

function changeChatRoom() {
	const newChat = document.getElementById("chatroom");
	if (newChat != null && newChat != selectedChat) {
		console.log(newChat);
	}
	return false;
}

function sendMessage() {
	const newMessage = document.getElementById("message");
	if (newMessage != null) {
		conn.send(newMessage.value);
	}
	return false;
}

window.onload = function () {
	document.getElementById("selectchat").onsubmit = changeChatRoom;
	document.getElementById("chatroommessage").onsubmit = sendMessage;
	if (window["WebSocket"]) {
		console.log("browser supports websockets");
		//connect to websocket here!!
		conn = new WebSocket("ws://" + document.location.host + "/ws");
		conn.onmessage = function (evt) {
			console.log(evt.data);
		};
	} else {
		alert("Browser does not support websockets");
	}
};
