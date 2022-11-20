import { useEffect, useState } from "react";

import Main from "./Main";
import About from "./components/About";
import Welcome from "./pages/Welcome";
import Home from "./pages/Home";
import OnboardingName from "./pages/OnboardingName";
import OnboardingAccess from "./pages/OnboardingAccess";
import ClusterLayout from "./pages/ClusterLayout";
import ClusterFiles from "./pages/ClusterFiles";
import ClusterMembers from "./pages/ClusterMembers";
import ClusterSettings from "./pages/ClusterSettings";
import Pinned from "./pages/Pinned";
import Shared from "./pages/Shared";
import Settings from "./pages/Settings";
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
          <Route path="/" element={<Main />}>
            <Route index element={<Home />} />
            <Route path="cluster" element={<ClusterLayout />}>
              <Route index element={<ClusterFiles />} />
              <Route path="members" element={<ClusterMembers />} />
              <Route path="settings" element={<ClusterSettings />} />
            </Route>
            <Route path="about" element={<About />} />
            <Route path="pinned" element={<Pinned />} />
            <Route path="shared" element={<Shared />} />
            <Route path="settings" element={<Settings />} />
          </Route>
        </Routes>
      ) : (
        <Routes>
          <Route index element={<Welcome />} />
          <Route path="onboarding-name" element={<OnboardingName />} />
          <Route path="onboarding-admins" element={<OnboardingAdmins />} />
          <Route path="onboarding-access" element={<OnboardingAccess />} />
          <Route path="onboarding-invite" element={<OnboardingInvite />} />
          <Route path="nft-selection" element={<NFTSelection />} />
        </Routes>
      )}
    </BrowserRouter>
  );
}

export default App;
