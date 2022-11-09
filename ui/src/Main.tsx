import Sidebar from "./components/Sidebar";
import Home from "./pages/Home";
import Pinned from "./pages/Pinned";

import { Routes, Route, Outlet } from "react-router-dom";

function Main() {
  return (
    <div>
      <Sidebar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="pinned" element={<Pinned />} />
      </Routes>
    </div>
  );
}

export default Main;
