import { useCallback } from "react"
import { Alchemy, Network } from "alchemy-sdk";
import { useAppStore } from "../store/app"

const useUser = () => {
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

  const setNfts = useAppStore((state) => state.setNfts);

  const getUserNfts = useCallback(async (userAddress: string) => {
    let allNfts: Array<Object> = [];
    const walletAddress: any = userAddress;

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
    setNfts(allNfts);
  }, [setNfts, polygonAlchemy.nft, mainnetAlchemy.nft])

  return {
    getUserNfts
  }
}

export default useUser;
