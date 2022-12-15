import { Link, Outlet } from 'react-router-dom';

function CreateLayout() {
  return (
    <div className="flex flex-col">
      <Link to=".." relative="path"><p className="text-l-slateblue-700">Back</p></Link>
      <div>
        <Outlet />
      </div>
    </div>
  )
}

export default CreateLayout;
