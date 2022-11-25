import { useCallback } from "react"
import axios from "axios"

const useClusters = () => {

  const getClusterMetadata = useCallback(async (clusterId: string) => {
    console.log('firing getClusterMetadata')
    var config = {
      method: 'get',
      url: `http://localhost:3000/api/cluster/${clusterId}`,
      headers: {}
    };

    axios(config)
      .then(function(response) {
        console.log(JSON.stringify(response.data));
      })
      .catch(function(error) {
        console.log(error);
      });
  }, [])

  return {
    getClusterMetadata
  }
}

export default useClusters;
