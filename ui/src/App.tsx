import { useEffect } from "react";

import { Files } from "./components/Files";
import Main from "./Main";
import { About } from "./components/About";
import Welcome from "./pages/Welcome";
import OnboardingName from "./pages/OnboardingName";
import OnboardingAccess from "./pages/OnboardingAccess";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { useEthers } from "@usedapp/core";
import OnboardingAdmins from "./pages/OnboardingAdmins";
import OnboardingInvite from "./pages/OnboardingInvite";
import NFTSelection from "./pages/NftSelection";

const client = new W3CWebSocket("ws://127.0.0.1:4000/api/ws");

function App() {
  const { account } = useEthers();

  useEffect(() => {
    console.log("starting websocket client");
    client.onopen = () => {
      client.send("Hello server!");
      console.log("WebSocket Client Connected");
    };
    client.onmessage = (message: any) => {
      console.log(message);
    };
  });

  return (
    <BrowserRouter>
      {account ? (
        <Routes>
          <Route path="/" element={<Main />} />
          <Route path="/about" element={<About />} />
        </Routes>
      ) : (
        <Routes>
          <Route path="/" element={<Welcome />} />
          <Route path="onboardingname" element={<OnboardingName />} />
          <Route path="onboardingadmins" element={<OnboardingAdmins />} />
          <Route path="onboardingaccess" element={<OnboardingAccess />} />
          <Route path="onboardinginvite" element={<OnboardingInvite />} />
          <Route path="nftselection" element={<NFTSelection />} />
        </Routes>
      )}
    </BrowserRouter>
  );
}

export default App;
