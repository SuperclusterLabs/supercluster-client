import Sidebar from "./components/Sidebar";

import { Outlet } from "react-router-dom";

function Main() {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <div className="bg-l-slateblue-300 w-screen p-12">
        <Outlet />
      </div>
    </div>
  );
}

export default Main;
