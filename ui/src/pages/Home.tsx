import { useState, useCallback, useRef, useMemo } from "react"
import { useAppStore } from "../store/app"
import TextInput from "../components/TextInput"
import ButtonPrimary from "../components/ButtonPrimary"
import useConversation from "../hooks/useConversation"

function Home() {
  const [address, setAddress] = useState<string>("0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20")

  const convoMessages = useAppStore((state) => state.convoMessages)
  const loadingConversations = useAppStore(
    (state) => state.loadingConversations
  )

  const messages = useMemo(
    () => convoMessages.get(address) ?? [],
    [convoMessages, address]
  )

  const messagesEndRef = useRef(null)

  const scrollToMessagesEndRef = useCallback(() => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ; (messagesEndRef.current as any)?.scrollIntoView({ behavior: 'smooth' })
  }, [])

  const { sendMessage } = useConversation(
    address,
    scrollToMessagesEndRef
  )

  const hasMessages = messages.length > 0

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setAddress(e.target.value)
  }

  async function handleSendMessage() {
    await sendMessage("test message")
  }

  if (loadingConversations && !hasMessages) {
    return (
      <div>
        <div className="flex">
          <h1 className="text-4xl font-bold text-[#111827]">👋 Loading conversations...</h1>
        </div>
      </div>
    )
  }

  return (
    <div>
      <div className="flex flex-col">
        <h1 className="text-4xl font-bold text-[#111827]">👋 Welcome to Supercluster Files!</h1>
        <TextInput value={address} placeholder="Address" onChange={handleInputChange} />
        <ButtonPrimary onClick={handleSendMessage} text="Send message" />
        {messages.map((message: any, i: number) => {
          return (
            <p key={i}>{message.content}</p>
          )
        })}
      </div>
    </div>
  )
}

export default Home;
