import { useEthers } from "@usedapp/core";

export function MetamaskConnect() {
  const { activateBrowserWallet } = useEthers();
  return (
    <div>
      <div>
        <button onClick={() => activateBrowserWallet()}>Connect</button>
      </div>
    </div>
  );
}
