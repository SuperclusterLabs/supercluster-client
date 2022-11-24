import Sidebar from "./components/Sidebar";

import { Outlet } from "react-router-dom";
import useListConversations from "./hooks/useListConversations"

function Main() {
  useListConversations();

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
