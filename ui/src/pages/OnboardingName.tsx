import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import superclusterLogo from "../assets/superclusterLogo.svg";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import { useEthers } from "@usedapp/core";
import { Alchemy, Network } from "alchemy-sdk";

function OnboardingName() {
  const navigate = useNavigate();
  const { account } = useEthers();
  const config = {
    apiKey: "98t_tAtPTdjvDoog8DbHxbpSRZgDAxv2",
    network: Network.ETH_MAINNET,
  };
  const alchemy = new Alchemy(config);

  const [clusterName, setClusterName] = useState<String>("");
  const [ens, setEns] = useState<String>("");
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    console.log("Onboarding Name triggered")
    localStorage.setItem("onboardingDone", "true");
    const getENS = async () => {
      const walletAddress: any = account; // replace with wallet address
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
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function confirmName() {
    console.log(clusterName);
    navigate("/onboardingadmins");
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setClusterName(e.target.value);
  }

  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        {!account && loading && <h1>Getting account</h1>}
        {!loading && (
          <div>
            <h1 className="text-4xl font-bold text-white mb-10">
              Hey 👋🏼, {ens === "" ? account : ens}! What should we name your
              cluster?
            </h1>
            <p className="text-2xl text-l-slategray-50">
              You'll need a name for your cluster. It will help your teammates
              find you a little easier. You can always change this afterwards.
            </p>
            <TextInput
              onChange={handleInputChange}
              placeholder="Cluster name"
            />
            <ButtonPrimary onClick={confirmName} text="Confirm name" />
          </div>
        )}
      </div>
    </div>
  );
}

export default OnboardingName;
