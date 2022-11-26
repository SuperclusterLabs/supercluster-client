import { useCallback } from "react"
import useInitXmtpClient from "../hooks/useInitXmtpClient";
import useWalletProvider from "../hooks/useWalletProvider";
import { useAppStore } from "../store/app";
import ButtonPrimary from "./ButtonPrimary";

export function MetamaskConnect() {
  const { initClient } = useInitXmtpClient()
  const signer = useAppStore((state) => state.signer)

  const { connect: connectWallet } = useWalletProvider();

  const handleLogin = useCallback(async (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    event.preventDefault()

    await connectWallet()
    signer && (await initClient(signer))
  }, [connectWallet, initClient, signer])

  return (
    <ButtonPrimary onClick={handleLogin} text="Connect wallet" />
  );
}
