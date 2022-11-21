import { useEffect, useState } from "react"
import ButtonPrimary from "../components/ButtonPrimary"
import TextInput from "../components/TextInput";

const exampleMembers = [
  {
    id: "id_1",
    clusters: ["exampleCluster", "exampleCluster"],
    ethAddr: "0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20",
    ipfsAddr: "ipfs_addr"
  },
  {
    id: "id_1",
    clusters: ["exampleCluster", "exampleCluster"],
    ethAddr: "0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20",
    ipfsAddr: "ipfs_addr"
  },
  {
    id: "id_1",
    clusters: ["exampleCluster", "exampleCluster"],
    ethAddr: "0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20",
    ipfsAddr: "ipfs_addr"
  },
]

function ClusterMembers() {
  // TODO: Get members from cluster API endpoint
  const [members, setMembers] = useState<Array<any>>([]);
  const [address, setAddress] = useState<string>("");

  // TODO: Set members from cluster API endpoint
  useEffect(() => {
    setMembers(exampleMembers)
  }, [])

  function handleInvite() {
    console.log("Inviting member:", address)
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setAddress(e.target.value);
  }

  return (
    <div className="flex flex-col mt-6 text-l-slateblue-700">
      <h2 className="font-bold text-xl">Current members</h2>
      <div className="flex flex-col space-y-3 mt-2 mb-9">
        {members.map((member: any) => {
          return (
            <p key={member.id}>{member.ethAddr}</p>
          )
        })}
      </div>
      <div>
        <h2 className="font-bold text-xl mb-4">Invite member</h2>
        <TextInput
          onChange={handleInputChange}
          placeholder="Enter address or ENS"
        />
        <ButtonPrimary onClick={handleInvite} text="Invite member" />
      </div>
    </div>
  )
}

export default ClusterMembers;
