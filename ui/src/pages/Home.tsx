import { useEffect } from "react"
import { useAppStore } from "../store/app"
import useClusters from "../hooks/useClusters";

function Home() {
  const clusterUserId = useAppStore((state) => state.clusterUserId)
  const setUserClusters = useAppStore((state) => state.setUserClusters)
  const { getUserClusters } = useClusters();

  useEffect(() => {
    const userClusters = async () => {
      if (clusterUserId) {
        await getUserClusters(clusterUserId)
      }
    }

    userClusters()
  }, [clusterUserId])

  return (
    <div>
      <div className="flex flex-col">
        <h1 className="text-4xl font-bold text-onboarding-bg">👋 Welcome to Supercluster Files!</h1>
      </div>
    </div>
  )
}

export default Home;
