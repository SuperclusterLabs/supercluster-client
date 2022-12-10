import create from 'zustand'
import { XMTPSlice, createXMTPSlice } from './xmtp'
import { WalletSlice, createWalletSlice } from "./wallet"
import { ClusterSlice, createClusterSlice } from "./cluster"

export const useAppStore = create<XMTPSlice & WalletSlice & ClusterSlice>()((...a) => ({
  ...createXMTPSlice(...a),
  ...createWalletSlice(...a),
  ...createClusterSlice(...a)
}))
