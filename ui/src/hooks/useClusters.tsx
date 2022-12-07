import { useCallback } from "react"
import axios from "axios"
import { useAppStore } from "../store/app"

const useClusters = () => {
  const setActiveCluster = useAppStore((state) => state.setActiveCluster)

  const getClusterMetadata = useCallback(async (clusterId: string) => {
    var config = {
      method: 'get',
      url: `http://localhost:3000/api/cluster/${clusterId}`,
      headers: {}
    };

    axios(config)
      .then(function (response) {
        console.log("Getting metadata:", response.data)
        setActiveCluster(response.data)
      })
      .catch(function (error) {
        console.log(error);
      });
  }, [setActiveCluster])

  return {
    getClusterMetadata
  }
}

export default useClusters;
