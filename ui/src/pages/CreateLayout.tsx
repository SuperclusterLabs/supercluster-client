import { Outlet } from 'react-router-dom';

function CreateLayout() {
  return (
    <div className="flex flex-col">
      <div className="flex items-center">
        <h1 className="text-4xl font-bold text-onboarding-bg">Create a new cluster</h1>
      </div>
      <div>
        <Outlet />
      </div>
    </div>
  )
}

export default CreateLayout;
