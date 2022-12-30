import { StateCreator } from 'zustand'
import { XMTPSlice } from './xmtp'
import { WalletSlice } from "./wallet"

interface userCluster {
  name: string,
  id: string,
}

interface createdCluster {
  name: string,
  id: string,
  nftAddr: string,
  admins: Array<string>,
}

export interface ClusterSlice {
  clusterUserId: string | undefined
  setClusterUserId: (clusterUserId: string | undefined) => void
  userClusters: Array<userCluster>
  setUserClusters: (userClusters: Array<userCluster>) => void
  createdCluster: createdCluster | undefined
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
  userClusters: [],
  setUserClusters: (userClusters: Array<userCluster>) => set(() => ({ userClusters })),
  createdCluster: undefined,
  setCreatedCluster: (createdCluster: any) => set(() => ({ createdCluster })),
})
