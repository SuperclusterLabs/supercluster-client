import { useEthers } from "@usedapp/core";

export function MetamaskConnect() {
  const { activateBrowserWallet } = useEthers();

  function login() {
    activateBrowserWallet();
  }

  return (
    <button
      className="bg-gradient-to-b from-l-success-main to-l-success-700 py-4 px-14 rounded-2xl"
      onClick={() => login()}
    >
      <span className="text-white font-bold text-md">Connect wallet</span>
    </button>
  );
}
