import React, { useState, useEffect, useRef } from "react";
import './App.css';

var ws: WebSocket;

function App() {
  const wsRef = useRef(ws);
  const [message, setMessage] = useState('');

  useEffect(() => {
    // ws.onopen = () => {
    //   console.log('connected')
    // }

    // ws.onmessage = (msg) => {
    //   console.log(msg)
    // }

    // ws.onclose = () => {
    //   console.log('closing connection')
    //   // ws.close()
    // }
  }, [])

  const onChange = (e: React.ChangeEvent<HTMLInputElement>)=> {
    setMessage(e.target.value);
  }

  return (
    <div className="App">
      <header className="App-header">
        <h2 style={{ color: "lavender" }}>Disgord</h2>
        <div className="input-group">
          <input 
            className="message"
            type="text"
            onChange={onChange}
            value={message}
          />
          <button className="btn" onClick={ join }>
              Join to server
          </button>
          <button className="btn" onClick={ () => send(message) }>
              Send
          </button>
        </div>
      </header>
    </div>
  );

  function join() {
    if (!wsRef.current) {
      ws = new WebSocket('ws://127.0.0.1:8080/join');

      ws.onopen = () => {
        console.log('connected')
      }
  
      ws.onmessage = msg => {
        console.log(msg)
      }
    }
    // connect();
  }

  function send(msg: string) {
    // ws.send('Hello, Server!');
    ws.send(msg);
  }
}

export default App;
