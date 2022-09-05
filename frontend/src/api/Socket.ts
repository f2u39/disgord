var ws = new WebSocket("ws://localhost:8080/join");

let connect = () => {
  console.log("Attempting Connection...");

  ws.onopen = () => {
    console.log("Successfully Connected");
  };

  ws.onmessage = msg => {
    console.log(msg);
  };

  ws.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  ws.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = (msg: string) => {
  console.log("sending msg: ", msg);
  ws.send(msg);
};

export { connect, sendMsg };