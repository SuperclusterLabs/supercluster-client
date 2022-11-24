import { useCallback } from "react"
import useInitXmtpClient from "../hooks/useInitXmtpClient";
import useWalletProvider from "../hooks/useWalletProvider";
import { useAppStore } from "../store/app";

export function MetamaskConnect() {
  const client = useAppStore((state) => state.client)
  const { initClient } = useInitXmtpClient()
  const walletAddress = useAppStore((state) => state.address)
  const signer = useAppStore((state) => state.signer)

  const { connect: connectWallet, disconnect: disconnectWallet } = useWalletProvider();

  const handleDisconnect = useCallback(async () => {
    await disconnectWallet()
  }, [disconnectWallet])

  const handleLogin = useCallback(async () => {
    await connectWallet()
    signer && (await initClient(signer))
  }, [connectWallet, initClient, signer])

  return (
    <button
      className="bg-gradient-to-b from-l-success-main to-l-success-700 py-4 px-14 rounded-2xl"
      onClick={() => handleLogin()}
    >
      <span className="text-white font-bold text-md">Connect wallet</span>
    </button>
  );
}
