import { Link, Outlet } from "react-router-dom"
import { useAppStore } from "../store/app"

function ClusterLayout() {
  const cluster = useAppStore((state) => state.activeCluster)

  if (!cluster) {
    return (
      <div className="flex flex-col">
        <div className="flex items-center">
          <h1 className="text-4xl font-bold text-[#111827]">Getting cluster...</h1>
        </div>
        <div>
          <Outlet />
        </div>
      </div>
    )
  }
  return (
    <div className="flex flex-col">
      <div className="flex items-center">
        <h1 className="text-4xl font-bold text-[#111827]">{cluster.name}</h1>
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
