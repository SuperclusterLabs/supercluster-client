import { useEffect, useState } from "react";

import Main from "./Main";
import About from "./components/About";
import Welcome from "./pages/Welcome";
import Home from "./pages/Home";
import OnboardingName from "./pages/OnboardingName";
import OnboardingAccess from "./pages/OnboardingAccess";
import Pinned from "./pages/Pinned";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { useEthers } from "@usedapp/core";
import OnboardingAdmins from "./pages/OnboardingAdmins";
import OnboardingInvite from "./pages/OnboardingInvite";
import NFTSelection from "./pages/NftSelection";

const client = new W3CWebSocket("ws://127.0.0.1:4000/api/ws");

function App() {
  const { account } = useEthers();

  const [onboardingDone, setOnboardingDone] = useState<boolean>(false);

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

  useEffect(() => {
    const onboardingDone = JSON.parse(
      localStorage.getItem("onboardingDone") || "false"
    );

    if (onboardingDone === "false") {
      setOnboardingDone(false);
      console.log("Local storage not found")
    } else {
      setOnboardingDone(true);
      console.log("Local Storage found")
    }
  }, []);

  return (
    <BrowserRouter>
      {account && onboardingDone ? (
        <Routes>
          <Route element={<Main />}>
            <Route index element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/pinned" element={<Pinned />} />
          </Route>
        </Routes>
      ) : (
        <Routes>
          <Route element={<Welcome />} />
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
