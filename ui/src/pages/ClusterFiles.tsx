import { useEffect, useState, useRef, useCallback, useMemo } from "react";
import Dropzone from "../components/Dropzone";
import { useAppStore } from "../store/app"
import axios from "axios"
import useConversation from "../hooks/useConversation"

function ClusterFiles() {
  const cluster = useAppStore((state) => state.activeCluster)
  const currentAddress = useAppStore((state) => state.address)
  const activeClusterNumberOfFiles = useAppStore((state) => state.activeClusterNumberOfFiles)
  const setActiveClusterNumberOfFiles = useAppStore((state) => state.setActiveClusterNumberOfFiles)

  // TODO: Need to get the files from the Cluster
  const [numberOfFiles, setNumberOfFiles] = useState<number>(0);

  const [address, setAddress] = useState<string>("")

  // TODO: Change to get Files from API
  useEffect(() => {
    if (cluster) {
      if (cluster.files) {
        setNumberOfFiles(cluster.files.length)
      }
    }

    if (currentAddress === "0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20") {
      setAddress("0xc45E269Bc5fe36B5b3D5934d4FF07BDD054787Ca")
    } else {
      setAddress("0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20")
    }
  }, [])

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
    setActiveClusterNumberOfFiles(2)
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
      console.log(e.target.files[0])
      formData.append('file', e.target.files[0])
      var config = {
        method: 'post',
        url: `http://localhost:3000/api/cluster/${cluster.id}`,
        data: formData
      };
      axios(config)
        .then(async (response: any) => {
          await sendMessage(response.data.file.id)
          setActiveClusterNumberOfFiles(activeClusterNumberOfFiles + 1)
        })
        .catch((error: any) => console.log(error))
    }
  }

  if (loadingConversations) {
    console.log("Getting message", messages)
  }

  return (
    <div>
      <div className="bg-white flex px-6 py-8 mt-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        <div className="text-center">
          <h2>Files pinned</h2>
          <p className="mt-2 font-bold text-3xl text-center">{activeClusterNumberOfFiles}</p>
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
          return (
            <div>{message.content}</div>
          )
        })}
      </div>
    </div>
  );
}

export default ClusterFiles;

