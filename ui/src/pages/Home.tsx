import { useState, useEffect } from "react";
import { useEthers } from "@usedapp/core";
import { Alchemy, Network } from "alchemy-sdk";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";

function Home() {
  const [ens, setEns] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(true);
  const [address, setAddress] = useState<string>("");

  const { account } = useEthers();
  const config = {
    apiKey: process.env.REACT_APP_ALCHEMY_MAINNET_API_KEY,
    network: Network.ETH_MAINNET,
  };
  const alchemy = new Alchemy(config);

  useEffect(() => {
    const getENS = async () => {
      const walletAddress: any = account;
      const ensContractAddress = "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85";
      const nfts = await alchemy.nft.getNftsForOwner(walletAddress, {
        contractAddresses: [ensContractAddress],
      });
      if (nfts.totalCount > 0) {
        setEns(nfts.ownedNfts[0].title);
      }
      setLoading(false);
    };
    getENS().catch(console.error);
  }, [account, alchemy.nft]);

  async function sendMessage() {
    console.log("sending a message")
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setAddress(e.target.value);
  }

  if (loading) {
    return (
      <div>
        <div className="flex">
          <h1 className="text-4xl font-bold text-onboarding-bg">🤔 Loading... </h1>
        </div>
      </div>
    )
  }

  if (ens !== "") {
    return (
      <div>
        <div className="flex flex-col">
          <h1 className="text-4xl font-bold text-onboarding-bg">👋 Welcome, {ens}!</h1>
          <TextInput value={address} placeholder="Recipient Address" onChange={handleInputChange} />
          <ButtonPrimary onClick={sendMessage} text="Send Message" />
        </div>
      </div>
    );
  } else {
    return (
      <div>
        <div className="flex">
          <h1 className="text-4xl font-bold text-onboarding-bg">👋 Welcome to Supercluster Files!</h1>
        </div>
      </div>
    )
  }
}

export default Home;
