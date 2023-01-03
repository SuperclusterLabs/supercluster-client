import { useCallback } from "react"
import axios from "axios"
import { useAppStore } from "../store/app"

const useClusters = () => {
  const setUserClusters = useAppStore((state) => state.setUserClusters);

  const getClusterMetadata = useCallback(async (clusterId: string) => {
    var config = {
      method: 'get',
      url: `http://localhost:3000/api/cluster/${clusterId}`,
      headers: {}
    };

    axios(config)
      .then(function(response) {
        console.log("Getting metadata:", response.data)
        return response.data
      })
      .catch(function(error) {
        console.log(error);
      });
  }, [])

  const getUserClusters = useCallback(async (userId: string) => {
    var config = {
      method: 'get',
      url: `http://localhost:3000/api/user/clusters?uId=${userId}`,
      headers: {}
    };

    axios(config)
      .then(function(response) {
        console.log("Getting user clusters:", response.data)
        if (response.data) {
          setUserClusters(response.data)
        }
      })
      .catch(function(error) {
        console.log(error);
      })
  }, [setUserClusters])

  return {
    getClusterMetadata,
    getUserClusters
  }
}

export default useClusters;
