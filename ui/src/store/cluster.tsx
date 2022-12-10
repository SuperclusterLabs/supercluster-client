import { StateCreator } from 'zustand'
import { XMTPSlice } from './xmtp'
import { WalletSlice } from "./wallet"

export interface ClusterSlice {
  clusterUserId: string | undefined
  setClusterUserId: (clusterUserId: string | undefined) => void
  userClusters: Array<string> | undefined | null
  setUserClusters: (userClusters: Array<string> | undefined | null) => void
  createdCluster: any | undefined
  setCreatedCluster: (createdCluster: any) => void
}

export const createClusterSlice: StateCreator<
  ClusterSlice & XMTPSlice & WalletSlice,
  [],
  [],
  ClusterSlice
> = (set) => ({
  clusterUserId: undefined,
  setClusterUserId: (clusterUserId: string | undefined) => set(() => ({ clusterUserId })),
  userClusters: undefined,
  setUserClusters: (userClusters: Array<string> | undefined | null) => set(() => ({ userClusters })),
  createdCluster: undefined,
  setCreatedCluster: (createdCluster: any) => set(() => ({ createdCluster })),
})
