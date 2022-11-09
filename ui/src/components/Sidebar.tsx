import { Link } from "react-router-dom";

function Sidebar() {
  return (
    <div>
      <h1>Sidebar</h1>
      <Link to="pinned">Pinned</Link>
    </div>
  );
}

export default Sidebar;
