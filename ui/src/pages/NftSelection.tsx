import { useState, useEffect } from "react";
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
    apiKey: process.env.REACT_APP_ALCHEMY_MAINNET_API_KEY,
    network: Network.ETH_MAINNET,
  };

  const polygonConfig = {
    apiKey: process.env.REACT_APP_ALCHEMY_MATIC_API_KEY,
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
    navigate("../onboarding-invite");
  }

  return (
    <div className="text-l-slateblue-700">
      <div>
        <h1 className="text-2xl font-bold">
          Alright! Which of your NFTs would you like to use for access
          control?
        </h1>
        <p className="text-lg">
          You can specify which NFTs your user need to own in order to access
          your cluster.
        </p>
        <div className="columns-5 my-8">
          {userNfts.map((nft: any, i: number) => {
            return (
              <div
                key={i}
                onClick={() => selectNft(nft)}
                className={`drop-shadow mb-4 p-4 rounded-2xl overflow-scroll cursor-pointer ${_.isEqual(nft, accessNft)
                  ? "text-white bg-l-slateblue-700"
                  : "text-l-slateblue-700 bg-white"
                  }`}
              >
                <h1 className="text-m">{nft.title}</h1>
              </div>
            );
          })}
        </div>
      </div>
      {accessNft ? (
        <ButtonPrimary onClick={confirmNft} text="Confirm NFT" />
      ) : null}
    </div>
  );
}

export default NFTSelection;
