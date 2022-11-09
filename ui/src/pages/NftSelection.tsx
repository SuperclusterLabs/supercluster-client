import { useState, useEffect } from "react";
import superclusterLogo from "../assets/superclusterLogo.svg";
import { Alchemy, Network } from "alchemy-sdk";
import { useEthers } from "@usedapp/core";
import _ from "underscore";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";

function NFTSelection() {
  const [userNfts, setUserNfts] = useState<Array<any>>([]);
  const [accessNft, setAccessNft] = useState<any>();

  const { account } = useEthers();
  const navigate = useNavigate();

  const mainnetConfig = {
    apiKey: "98t_tAtPTdjvDoog8DbHxbpSRZgDAxv2",
    network: Network.ETH_MAINNET,
  };

  const polygonConfig = {
    apiKey: "hKo6uza_jM6M9SI9MhumJJEU1qlswEXx",
    network: Network.MATIC_MAINNET,
  };

  const mainnetAlchemy = new Alchemy(mainnetConfig);
  const polygonAlchemy = new Alchemy(polygonConfig);

  useEffect(() => {
    const getNfts = async () => {
      let allNfts: Array<Object> = [];
      const walletAddress: any = account;
      let userMainnetNfts = await mainnetAlchemy.nft.getNftsForOwner(
        walletAddress
      );
      let userPolygonNfts = await polygonAlchemy.nft.getNftsForOwner(
        walletAddress
      );
      if (userMainnetNfts.ownedNfts.length > 0) {
        allNfts = allNfts.concat(userMainnetNfts.ownedNfts);
      }
      if (userPolygonNfts.ownedNfts.length > 0) {
        allNfts = allNfts.concat(userPolygonNfts.ownedNfts);
      }
      console.log(allNfts);
      setUserNfts(allNfts);
    };

    getNfts();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function selectNft(nft: any) {
    setAccessNft(nft);
  }

  function confirmNft() {
    navigate("/onboardinginvite");
  }

  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <div>
          <h1 className="text-4xl font-bold text-white mb-10">
            Alright! Which of your NFTs would you like to use for access
            control?
          </h1>
          <p className="text-2xl text-l-slategray-50">
            You can specify which NFTs your user need to own in order to access
            your cluster.
          </p>
          <div className="columns-5 mt-11">
            {userNfts.map((nft: any, i: number) => {
              return (
                <div
                  key={i}
                  onClick={() => selectNft(nft)}
                  className={`mb-4 p-4 rounded-2xl overflow-scroll cursor-pointer ${
                    _.isEqual(nft, accessNft)
                      ? "text-l-slateblue-primary bg-l-slateblue-700"
                      : "text-l-slateblue-700 bg-l-slateblue-primary"
                  }`}
                >
                  <h1 className="font-bold text-m">{nft.title}</h1>
                </div>
              );
            })}
          </div>
        </div>
        {accessNft ? (
          <ButtonPrimary onClick={confirmNft} text="Confirm NFT" />
        ) : null}
      </div>
    </div>
  );
}

export default NFTSelection;