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
        let activeCluster = await getClusterMetadata(userClusters[0])
        console.log("active cluster:", activeCluster)
        setActiveCluster(activeCluster)
      }
    }
    getActiveCluster()
  }, [userClusters])

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
