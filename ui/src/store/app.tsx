import { Client, Conversation, DecodedMessage } from '@xmtp/xmtp-js'
import { Signer } from 'ethers'
import create from 'zustand'

interface AppState {
  signer: Signer | undefined
  setSigner: (signer: Signer | undefined) => void
  address: string | undefined
  setAddress: (address: string | undefined) => void
  client: Client | undefined | null
  setClient: (client: Client | undefined | null) => void
  conversations: Map<string, Conversation>
  setConversations: (conversations: Map<string, Conversation>) => void
  loadingConversations: boolean
  setLoadingConversations: (loadingConversations: boolean) => void
  convoMessages: Map<string, DecodedMessage[]>
  setConvoMessages: (value: Map<string, DecodedMessage[]>) => void
  clusterUserId: string | undefined
  setClusterUserId: (clusterUserId: string | undefined) => void
  userClusters: Array<string> | undefined | null
  setUserClusters: (userClusters: Array<string> | undefined | null) => void
  createdCluster: any | undefined
  setCreatedCluster: (createdCluster: any) => void
}

export const useAppStore = create<AppState>((set) => ({
  signer: undefined,
  setSigner: (signer: Signer | undefined) => set(() => ({ signer })),
  address: undefined,
  setAddress: (address: string | undefined) => set(() => ({ address })),
  client: undefined,
  setClient: (client: Client | undefined | null) => set(() => ({ client })),
  conversations: new Map(),
  setConversations: (conversations: Map<string, Conversation>) =>
    set(() => ({ conversations })),
  loadingConversations: false,
  setLoadingConversations: (loadingConversations: boolean) =>
    set(() => ({ loadingConversations })),
  convoMessages: new Map(),
  setConvoMessages: (convoMessages: Map<string, DecodedMessage[]>) =>
    set(() => ({ convoMessages })),
  clusterUserId: undefined,
  setClusterUserId: (clusterUserId: string | undefined) => set(() => ({ clusterUserId })),
  userClusters: undefined,
  setUserClusters: (userClusters: Array<string> | undefined | null) => set(() => ({ userClusters })),
  createdCluster: undefined,
  setCreatedCluster: (createdCluster: any) => set(() => ({ createdCluster })),
}))
