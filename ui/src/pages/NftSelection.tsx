import { useState, useEffect } from "react";
import _ from "underscore";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";
import { useAppStore } from "../store/app"
import useUser from "../hooks/useUser";
import axios from "axios";

function NFTSelection() {
  const { getUserNfts } = useUser();
  const [accessNft, setAccessNft] = useState<any>();

  const createdCluster = useAppStore((state) => state.createdCluster)
  const address = useAppStore((state) => state.address)
  const nfts = useAppStore((state) => state.nfts)

  const navigate = useNavigate();

  useEffect(() => {
    if (address) {
      getUserNfts(address)
    }
  }, [address, getUserNfts]);

  function selectNft(nft: any) {
    setAccessNft(nft);
  }

  async function confirmNft() {
    if (createdCluster) {
      let data = createdCluster
      data.nftAddr = accessNft.contract.address
      let config = {
        method: 'put',
        url: `http://localhost:3000/api/cluster/${data.id}`,
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
  }

  function renderImage(metadata: any) {
    if (metadata.image) {
      return metadata.image
    } else if (metadata.image_url) {
      return metadata.image_url
    }
    return "";
  }

  return (
    <div className="text-l-slateblue-700 mt-6">
      <div>
        <h1 className="text-4xl font-bold">
          Alright! Which of your NFTs would you like to use for access
          control?
        </h1>
        <p className="text-xl mt-4">
          You can specify which NFTs your user need to own in order to access
          your cluster.
        </p>
        <div className="columns-4 gap-y-4 my-8">
          {nfts.map((nft: any, i: number) => {
            return (
              <div
                key={i}
                onClick={() => selectNft(nft)}
                className={`flex flex-col p-4 drop-shadow rounded-2xl cursor-pointer ${_.isEqual(nft, accessNft)
                  ? "text-white bg-l-slateblue-700"
                  : "text-l-slateblue-700 bg-white"
                  }`}
              >
                <img alt="Base profile of the NFT" src={renderImage(nft.rawMetadata)} />
                <h1 className="text-l mt-6">{nft.title}</h1>
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
