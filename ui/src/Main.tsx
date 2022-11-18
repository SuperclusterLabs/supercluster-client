import Sidebar from "./components/Sidebar";

import { Outlet } from "react-router-dom";

function Main() {
  return (
    <div>
      <Sidebar />
      <Outlet />
    </div>
  );
}

export default Main;
