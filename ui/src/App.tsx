import { useEffect } from "react";

// import { Files } from "./components/Files";
import { About } from "./components/About";
import Welcome from "./pages/Welcome";
import OnboardingName from "./pages/OnboardingName";
import OnboardingAccess from "./pages/OnboardingAccess";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { w3cwebsocket as W3CWebSocket } from "websocket";
// import { useEthers } from "@usedapp/core";
import OnboardingAdmins from "./pages/OnboardingAdmins";
import OnboardingInvite from "./pages/OnboardingInvite";
import NFTSelection from "./pages/NftSelection";

const client = new W3CWebSocket("ws://127.0.0.1:4000/api/ws");

function App() {
  // const { account } = useEthers();

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
    console.log(process.env);
  });

  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            // account ? (
            //   <div className="container">
            //     <Files />
            //   </div>
            // ) : (
            //   <Welcome />
            // )
            <Welcome />
          }
        />
        <Route path="/onboarding-name" element={<OnboardingName />} />
        <Route path="/onboarding-admins" element={<OnboardingAdmins />} />
        <Route path="/onboarding-access" element={<OnboardingAccess />} />
        <Route path="/onboarding-invite" element={<OnboardingInvite />} />
        <Route path="/nft-selection" element={<NFTSelection />} />
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

export default App;
