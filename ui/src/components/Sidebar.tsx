import { Link } from "react-router-dom"
import superclusterLogo from "../assets/superclusterLogo.svg"
import { useAppStore } from "../store/app"

function Sidebar() {
  const userClusters = useAppStore((state) => state.userClusters)

  return (
    <nav className="flex flex-col bg-onboarding-bg text-l-slateblue-primary pt-6 px-9">
      <Link className="mb-10" to="/">
        <img
          className="max-w-none h-[16px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
      </Link>
      <div className="flex flex-col space-y-6">
        <p>🪐 Clusters</p>
        {userClusters.length ? (
          <ul className="pl-5">
            {userClusters.map((cluster: any) => (
              <li className="py-2" key={cluster.id}>
                <Link to={`cluster/${cluster.id}`}>
                  {cluster.name}
                </Link>
              </li>
            ))}
          </ul>
        ) : null}
        <Link className="pl-5" to="create">Create cluster</Link>
        <Link to="pinned">📌 Pinned</Link>
        <Link to="shared">📁 Shared</Link>
        <Link to="settings">🧰 Settings</Link>
      </div>
    </nav>
  );
}

export default Sidebar;
