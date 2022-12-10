import { useState, useRef, useCallback, useMemo } from "react";
import Dropzone from "../components/Dropzone";
import { useAppStore } from "../store/app"
import axios from "axios"
import useConversation from "../hooks/useConversation"

function ClusterFiles() {
  const currentAddress = useAppStore((state) => state.address)

  const [address, setAddress] = useState<string>("")

  const convoMessages = useAppStore((state) => state.convoMessages)

  const messages = useMemo(
    () => convoMessages.get(address) ?? [],
    [convoMessages, address]
  )

  const messagesEndRef = useRef(null)

  // TODO: Rename this function. This is the callback that happens when a new message
  // is received from a channel.
  const scrollToMessagesEndRef = useCallback(() => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    ; (messagesEndRef.current as any)?.scrollIntoView({ behavior: 'smooth' })
  }, [])

  const { sendMessage } = useConversation(
    address,
    scrollToMessagesEndRef
  )

  function handleFileUpload(e: React.ChangeEvent<HTMLInputElement>) {
    const formData = new FormData();

    if (e.target.files) {
      formData.append('file', e.target.files[0])
      // var config = {
      //   method: 'post',
      //   url: `http://localhost:3000/api/cluster/${cluster.id}`,
      //   data: formData
      // };
      // axios(config)
      //   .then(async (response: any) => {
      //     await sendMessage(response.data.file.id)
      //   })
      //   .catch((error: any) => console.log(error))
    }
  }

  return (
    <div>
      <div className="bg-white flex px-6 py-8 mt-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        <div className="text-center">
          <h2>Files pinned</h2>
        </div>
        <div className="text-center">
          <h2>Total members</h2>
          <p className="mt-2 font-bold text-3xl">2</p>
        </div>
      </div>
      <div className="flex items-center mt-6">
        <h2 className="font-bold text-3xl mr-6">Files</h2>
      </div>
      <div className="flex mt-4">
        <Dropzone multiple={true} onChange={handleFileUpload} />
      </div>
      <div className="mt-8 space-y-3">
        {messages.map((message: any) => {
          // TODO: Create a table for the files
          return (
            <div>{message.content}</div>
          )
        })}
      </div>
    </div>
  );
}

export default ClusterFiles;
