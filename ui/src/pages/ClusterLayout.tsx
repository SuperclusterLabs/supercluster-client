import { Link, Outlet } from "react-router-dom"
import { useAppStore } from "../store/app"

function ClusterLayout() {
  return (
    <div className="flex flex-col">
      <div className="flex items-center">
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
