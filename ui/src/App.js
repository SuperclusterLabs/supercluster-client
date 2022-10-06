import React from 'react';

import { Files } from "./components/Files";
import { About } from "./components/About";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";

const client = new W3CWebSocket('ws://127.0.0.1:4000/api/ws');

class App extends React.Component {
  componentWillMount() {
    console.log("starting websocket client")
    client.onopen = () => {
      client.send("Hello server!");
      console.log('WebSocket Client Connected');
    };
    client.onmessage = (message) => {
      console.log(message);
    };
  }

  render() {
    return (
      <BrowserRouter>
        <Routes>
          <Route
            path="/"
            element={
              <div className="container">
                <Files />
              </div>
            }
          />
          <Route
            path="/about"
            element={
              <div className="container">
                <About />
              </div>
            }
          />
        </Routes>
      </BrowserRouter>
    );
  }
}

export default App;
