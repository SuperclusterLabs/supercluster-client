import { useState, useEffect } from "react";
import { Alchemy, Network } from "alchemy-sdk";
import _ from "underscore";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";
import { useAppStore } from "../store/app"
import axios from "axios";

function NFTSelection() {
  const [userNfts, setUserNfts] = useState<Array<any>>([]);
  const [accessNft, setAccessNft] = useState<any>();
  const createdCluster = useAppStore((state) => state.createdCluster)

  const address = useAppStore((state) => state.address)

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
      const walletAddress: any = address;
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

  async function confirmNft() {
    console.log(accessNft);
    let data = createdCluster
    data.nftAddr = accessNft.contract.address
    let config = {
      method: 'put',
      url: `http://localhost:3000/api/cluster/${createdCluster.id}`,
      headers: {
        'Content-Type': 'text/plain'
      },
      data: data
    };

    await axios(config)
      .then(function(response) {
        console.log(JSON.stringify(response.data));
        navigate("../onboarding-invite");
      })
      .catch(function(error) {
        console.log(error);
      });
  }

  return (
    <div className="text-[#334574] mt-6">
      <div>
        <h1 className="text-2xl font-bold mt-4">
          Alright! Which of your NFTs would you like to use for access
          control?
        </h1>
        <p className="text-lg mt-4">
          You can specify which NFTs your user need to own in order to access
          your cluster.
        </p>
        <div className="columns-5 my-8">
          {userNfts.map((nft: any, i: number) => {
            return (
              <div
                key={i}
                onClick={() => selectNft(nft)}
                className={`drop-shadow mb-4 p-4 rounded-2xl cursor-pointer ${_.isEqual(nft, accessNft)
                  ? "text-white bg-[#334574]"
                  : "text-[#334574] bg-white"
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
