import Main from "./Main";
import Welcome from "./pages/Welcome";
import Home from "./pages/Home";
import OnboardingName from "./pages/OnboardingName";
import OnboardingAccess from "./pages/OnboardingAccess";
import ClusterLayout from "./pages/ClusterLayout";
import ClusterFiles from "./pages/ClusterFiles";
import ClusterMembers from "./pages/ClusterMembers";
import ClusterSettings from "./pages/ClusterSettings";
import Pinned from "./pages/Pinned";
import Shared from "./pages/Shared";
import Settings from "./pages/Settings";
import NotFound from "./pages/NotFound";
import { createBrowserRouter, RouterProvider, BrowserRouter, Routes, Route } from "react-router-dom";
import CreateLayout from "./pages/CreateLayout";
import OnboardingAdmins from "./pages/OnboardingAdmins";
import OnboardingInvite from "./pages/OnboardingInvite";
import NFTSelection from "./pages/NftSelection";
import AddressSelection from "./pages/AddressSelection";
import { useAppStore } from "./store/app"

const mainRouter = createBrowserRouter([
  {
    path: "/",
    element: <Main />,
    errorElement: <NotFound />,
    children: [
      {
        path: "cluster"
      }
    ]
  }
])

const welcomeRouter = createBrowserRouter([
  {
    path: "/",
    element: <Welcome />
  }
])

function App() {
  const address = useAppStore((state) => state.address)
  const userClusters = useAppStore((state) => state.userClusters)

  if (address) {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Main />}>
            <Route index element={<Home />} />
            <Route path="cluster" element={<ClusterLayout />}>
              <Route index element={<ClusterFiles />} />
              <Route path="members" element={<ClusterMembers />} />
              <Route path="settings" element={<ClusterSettings />} />
            </Route>
            <Route path="pinned" element={<Pinned />} />
            <Route path="shared" element={<Shared />} />
            <Route path="settings" element={<Settings />} />
            <Route path="create" element={<CreateLayout />}>
              <Route index element={<OnboardingName />} />
              <Route path="onboarding-admins" element={<OnboardingAdmins />} />
              <Route path="onboarding-access" element={<OnboardingAccess />} />
              <Route path="nft-selection" element={<NFTSelection />} />
              <Route path="address-selection" element={<AddressSelection />} />
              <Route path="onboarding-invite" element={<OnboardingInvite />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    )
  }

  return (
    <RouterProvider router={welcomeRouter} />
  );
}

export default App;
