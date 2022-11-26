import { useEffect } from "react"
import Sidebar from "./components/Sidebar";
import { useAppStore } from "./store/app"
import { Outlet } from "react-router-dom";
import useClusters from "./hooks/useClusters"

function Main() {
  const userClusters = useAppStore((state) => state.userClusters)
  const setActiveCluster = useAppStore((state) => state.setActiveCluster)
  const { getClusterMetadata } = useClusters()

  useEffect(() => {
    const getActiveCluster = async () => {
      if (userClusters && userClusters !== undefined) {
        console.log('user clusters:', userClusters)
        let activeCluster = await getClusterMetadata(userClusters[0])
        console.log("active cluster:", activeCluster)
        setActiveCluster(activeCluster)
      }
    }
    getActiveCluster()
  }, [userClusters, getClusterMetadata, setActiveCluster])

  return (
    <div className="flex h-screen">
      <Sidebar />
      <div className="bg-[#F8FAFC] w-screen p-12">
        <Outlet />
      </div>
    </div>
  );
}

export default Main;
