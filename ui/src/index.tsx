import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import { Mainnet, DAppProvider } from "@usedapp/core";
import { getDefaultProvider } from "ethers";
import reportWebVitals from "./reportWebVitals";
import { store } from "./store/store"
import { Provider } from "react-redux"

const config = {
  readOnlyChainId: Mainnet.chainId,
  readOnlyUrls: {
    [Mainnet.chainId]: getDefaultProvider("mainnet"),
  },
};

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <DAppProvider config={config}>
        <App />
      </DAppProvider>
    </Provider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
