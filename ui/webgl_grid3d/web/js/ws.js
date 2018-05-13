var walker = require('../index.js');

var ws = new WebSocket('ws://' + window.location.host + '/ws');

ws.onopen = function (event) {
	ws.send('{"cmd": "init"}'); 
};

ws.onmessage = function (event) {
	let msg = JSON.parse(event.data);
	switch(msg.type) {
		case "data":
			walker.addData(msg.data);
			break;
	}
}

module.exports = { ws };
