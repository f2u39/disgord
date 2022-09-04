import React, { useState } from "react";
import './App.css';

function App() {
  const ws = new WebSocket('ws://127.0.0.1:8080/join');


  // ws.onopen = () => {
  //   // ws.send("Hello server!");
  //   console.log('WebSocket Client Connected');
  // };
  // ws.onmessage = (message) => {
  //   console.log(message);
  // };

  return (
    <div className="App">
      <header className="App-header">
        <h2 style={{ color: "lavender" }}>👾Disgord</h2>
        <form>
          <input className="name" />
          <button className="btn" onClick={ join }>
              Join to server
          </button>
          <button className="btn" onClick={ send }>
              Send
          </button>
        </form>
      </header>
    </div>
  );

  function join() {  
    ws.send('Hello, Server!');
  }

  function send() {
    ws.send('Hello, Server!');
  }
}

export default App;
