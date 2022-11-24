import { Link, Outlet } from "react-router-dom"

function ClusterLayout() {
  return (
    <div className="flex flex-col">
      <div className="flex items-center">
        <h1 className="text-4xl font-bold text-[#111827]">BanklessDAO</h1>
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
