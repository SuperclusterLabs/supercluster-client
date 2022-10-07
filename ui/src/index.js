import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./App";
import { Mainnet, DAppProvider } from '@usedapp/core'
import { getDefaultProvider } from 'ethers'

const config = {
  readOnlyChainId: Mainnet.chainId,
  readOnlyUrls: {
    [Mainnet.chainId]: getDefaultProvider('mainnet'),
  },
}

ReactDOM.render(
  <React.StrictMode>
    <DAppProvider config={config}>
      <App />
    </DAppProvider>
  </React.StrictMode>,
  document.getElementById("root")
);
