import { useCallback } from 'react';
import { ethers } from 'ethers';
import { useAppStore } from "../store/app"
import { useEthers } from '@usedapp/core';

let provider: ethers.providers.Web3Provider


const useWalletProvider = () => {
  const setAddress = useAppStore((state) => state.setAddress)
  const setSigner = useAppStore((state) => state.setSigner)
  const { activateBrowserWallet, account } = useEthers()

  const connect = useCallback(async () => {
    activateBrowserWallet();

    try {
      provider = new ethers.providers.Web3Provider(window.ethereum, 'any')
      const newSigner = provider.getSigner()
      setSigner(newSigner)
      setAddress(await newSigner.getAddress())
      return newSigner
    } catch (e) {
      console.log("error:", e)
    }
  }, [activateBrowserWallet, setSigner, setAddress])

  const disconnect = useCallback(() => {
    if (!account) return
    localStorage.removeItem('walletconnect')
    Object.keys(localStorage).forEach((key) => {
      if (key.startsWith('-walletlink')) {
        localStorage.removeItem(key)
      }
    })
    setSigner(undefined)
    setAddress(undefined)
  }, [account, setAddress, setSigner])

  return {
    connect,
    disconnect
  }
}

export default useWalletProvider;
