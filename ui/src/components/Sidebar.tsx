import { Link } from "react-router-dom";
import superclusterLogo from "../assets/superclusterLogo.svg";

function Sidebar() {
  return (
    <div className="flex flex-col bg-onboarding-bg text-r-slateblue-primary pt-6 px-9">
      <Link className="mb-10" to="/">
        <img
          className="max-w-none h-[16px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
      </Link>
      <div className="flex flex-col space-y-6">
        <Link to="cluster">🪐 Clusters</Link>
        <Link className="pl-5" to="create">Create cluster</Link>
        <Link to="pinned">📌 Pinned</Link>
        <Link to="shared">📁 Shared</Link>
        <Link to="settings">🧰 Settings</Link>
      </div>
    </div>
  );
}

export default Sidebar;
