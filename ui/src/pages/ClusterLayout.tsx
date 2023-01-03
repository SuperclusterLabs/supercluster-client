import { useEffect, useState } from "react"
import { Link, Outlet } from "react-router-dom"
import { useAppStore } from "../store/app";
import { useParams } from "react-router-dom";
import _ from "underscore";

function ClusterLayout() {
  const { clusterId } = useParams();
  const userClusters = useAppStore((state) => state.userClusters);
  const [clusterName, setClusterName] = useState<string>("");

  useEffect(() => {
    const cluster = _.where(userClusters, { id: clusterId })
    if (cluster) {
      console.log(cluster)
      setClusterName(cluster[0].name)
    }
  }, [clusterId, userClusters])

  return (
    <div className="flex flex-col">
      <div className="flex items-center">
        <h1 className="text-l-slateblue-700 font-bold text-4xl">{clusterName}</h1>
        <div className="ml-10 space-x-8">
          <Link to="/cluster">Files</Link>
          <Link to="/cluster/members">Members</Link>
          <Link to="/cluster/settings">Settings</Link>
        </div>
      </div>
      <div>
        <Outlet />
      </div>
    </div>
  )
}

export default ClusterLayout
