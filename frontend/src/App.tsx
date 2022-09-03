import React from 'react';
import ReactDOM from 'react-dom';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h2 style={{ color: "lavender" }}>👾Disgord</h2>
        <form>
          <input className="name" />
          <button className="join" onClick={ join }>
              Join to server
          </button>
        </form>
        
      </header>
    </div>
  );
}

function join() {
  var uri = 'ws://127.0.0.1:8080/join';
  var ws = new WebSocket(uri)

  ws.onopen = function() {
    console.log('Connected')
  }

  setInterval(function() {
    ws.send('Hello, Server!');
  }, 1000);
}

export default App;
