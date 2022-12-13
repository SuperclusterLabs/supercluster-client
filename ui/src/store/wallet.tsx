import { Signer } from 'ethers'
import { StateCreator } from 'zustand'
import { XMTPSlice } from "./xmtp"
import { ClusterSlice } from "./cluster"

export interface WalletSlice {
  signer: Signer | undefined
  setSigner: (signer: Signer | undefined) => void
  address: string | undefined
  setAddress: (address: string | undefined) => void
}

export const createWalletSlice: StateCreator<
  WalletSlice & XMTPSlice & ClusterSlice,
  [],
  [],
  WalletSlice
> = (set) => ({
  signer: undefined,
  setSigner: (signer: Signer | undefined) => set(() => ({ signer })),
  address: undefined,
  setAddress: (address: string | undefined) => set(() => ({ address })),
})
