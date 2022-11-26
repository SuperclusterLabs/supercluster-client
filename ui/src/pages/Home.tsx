import { useEffect } from "react"
import { useAppStore } from "../store/app"
import useClusters from "../hooks/useClusters"

function Home() {
  const userClusters = useAppStore((state) => state.userClusters)
  const setActiveCluster = useAppStore((state) => state.setActiveCluster)

  const { getClusterMetadata } = useClusters()

  useEffect(() => {
    if (userClusters && userClusters !== undefined) {
      let activeCluster = getClusterMetadata(userClusters[0])
      setActiveCluster(activeCluster)
    }
  }, [userClusters, getClusterMetadata])
  return (
    <div>
      <div className="flex flex-col">
        <h1 className="text-4xl font-bold text-[#111827]">👋 Welcome to Supercluster Files!</h1>
      </div>
    </div>
  )
}

export default Home;
