import { Client, Conversation, DecodedMessage } from '@xmtp/xmtp-js'
import { StateCreator } from 'zustand'
import { WalletSlice } from './wallet'
import { ClusterSlice } from "./cluster"

export interface XMTPSlice {
  client: Client | undefined | null
  setClient: (client: Client | undefined | null) => void
  conversations: Map<string, Conversation>
  setConversations: (conversations: Map<string, Conversation>) => void
  loadingConversations: boolean
  setLoadingConversations: (loadingConversations: boolean) => void
  convoMessages: Map<string, DecodedMessage[]>
  setConvoMessages: (value: Map<string, DecodedMessage[]>) => void
}

export const createXMTPSlice: StateCreator<
  XMTPSlice & WalletSlice & ClusterSlice,
  [],
  [],
  XMTPSlice
> = (set) => ({
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
})
