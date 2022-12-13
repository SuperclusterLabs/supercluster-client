import { useCallback, useEffect, useState } from 'react';
import { ethers } from 'ethers';
import { useAppStore } from "../store/app"
import Web3Modal, { IProviderOptions, providers } from 'web3modal';
import WalletConnectProvider from '@walletconnect/web3-provider';
import axios from "axios";

// Ethereum mainnet
const ETH_CHAIN_ID = 1

const cachedLookupAddress = new Map<string, string | undefined>()
const cachedResolveName = new Map<string, string | undefined>()

let provider: ethers.providers.Web3Provider

const useWalletProvider = () => {
  const setAddress = useAppStore((state) => state.setAddress)
  const setSigner = useAppStore((state) => state.setSigner)
  const setClusterUserId = useAppStore((state) => state.setClusterUserId)
  const setUserClusters = useAppStore((state) => state.setUserClusters)

  const [web3Modal, setWeb3Modal] = useState<Web3Modal>()

  // Set the web3 connector modal
  useEffect(() => {
    const infuraId = process.env.REACT_INFURA_ID || '05b704a736774e568be9c6ad3f989c9c'
    const providerOptions: IProviderOptions = {
      walletconnect: {
        package: WalletConnectProvider,
        options: {
          infuraId
        }
      }
    }
    // Check if the user has MetaMask. If not, take them to MetaMask download page
    if (!window.ethereum || !window.ethereum.isMetaMask) {
      providerOptions['custom-metamask'] = {
        display: {
          logo: providers.METAMASK.logo,
          name: "Install MetaMask",
          description: 'Connect using browser wallet'
        },
        package: {},
        connector: async () => {
          window.open('https://metamask.io')
        }
      }
    }
    setWeb3Modal(new Web3Modal({ cacheProvider: true, providerOptions }))
  }, [])

  // Look to see if there's a provider that's already been cached in local storage
  useEffect(() => {
    if (!web3Modal) return
    const initCached = async () => {
      const cachedProviderJson = localStorage.getItem(
        'WEB3_CONNECT_CACHED_PROVIDER'
      )
      if (!cachedProviderJson) return
      const cachedProviderName = JSON.parse(cachedProviderJson)
      const instance = await web3Modal.connectTo(cachedProviderName)
      if (!instance) return
      provider = new ethers.providers.Web3Provider(instance, 'any')
      const newSigner = provider.getSigner()
      setSigner(newSigner)
      setAddress(await newSigner.getAddress())

      let config = {
        method: 'post',
        url: 'http://localhost:3000/api/user',
        headers: {
          'Content-Type': 'text/plain'
        },
        data: { "ethAddr": await newSigner.getAddress() }
      };

      axios(config)
        .then(function(response: any) {
          setClusterUserId(response.data.id);
        })
        .catch(function(error) {
          console.error(error);
        });
    }
    initCached()
  }, [web3Modal, setSigner, setAddress, setClusterUserId, setUserClusters])

  const resolveName = useCallback(async (name: string) => {
    if (cachedResolveName.has(name)) {
      return cachedResolveName.get(name)
    }

    const { chainId } = (await provider?.getNetwork()) || {}

    if (Number(chainId) !== ETH_CHAIN_ID) {
      return undefined
    }
    const address = (await provider?.resolveName(name)) || undefined
    cachedResolveName.set(name, address)
    return address
  }, [])

  const lookupAddress = useCallback(async (address: string) => {
    if (cachedLookupAddress.has(address)) {
      return cachedLookupAddress.get(address)
    }
    const { chainId } = (await provider?.getNetwork()) || {}

    if (Number(chainId) !== ETH_CHAIN_ID) {
      return undefined
    }

    const name = (await provider?.lookupAddress(address)) || undefined
    cachedLookupAddress.set(address, name)
    return name
  }, [])

  const connect = useCallback(async () => {
    if (!web3Modal) throw new Error('web3Modal not initialized')
    try {
      const instance = await web3Modal.connect()
      if (!instance) return
      provider = new ethers.providers.Web3Provider(instance, 'any')
      const newSigner = provider.getSigner()
      setSigner(newSigner)
      setAddress(await newSigner.getAddress())

      let config = {
        method: 'post',
        url: 'http://localhost:3000/api/user',
        headers: {
          'Content-Type': 'text/plain'
        },
        data: { "ethAddr": await newSigner.getAddress() }
      };

      axios(config)
        .then(function(response: any) {
          setClusterUserId(response.data.id);
          if (response.clusters !== null) {
            setUserClusters(response.data.clusters)
          }
        })
        .catch(function(error) {
          console.log(error);
        });

      return newSigner
    } catch (e) {
      // TODO: better error handling/surfacing here.
      // Note that web3Modal.connect throws an error when the user closes the
      // modal, as "User closed modal"
      console.log('error', e)
    }
  }, [web3Modal, setSigner, setAddress, setClusterUserId, setUserClusters])

  const disconnect = useCallback(() => {
    if (!web3Modal) return
    web3Modal.clearCachedProvider()
    localStorage.removeItem('walletconnect')
    Object.keys(localStorage).forEach((key) => {
      if (key.startsWith('-walletlink')) {
        localStorage.removeItem(key)
      }
    })
    setSigner(undefined)
    setAddress(undefined)
  }, [web3Modal, setSigner, setAddress])

  return {
    resolveName,
    lookupAddress,
    connect,
    disconnect,
  }
}

export default useWalletProvider;
