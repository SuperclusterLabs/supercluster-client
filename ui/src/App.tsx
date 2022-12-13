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
import { createBrowserRouter, RouterProvider } from "react-router-dom";
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
        element: <Home />,
        index: true
      },
      {
        path: "cluster/:clusterId",
        element: <ClusterLayout />,
        errorElement: <NotFound />,
        children: [
          {
            element: <ClusterFiles />,
            index: true
          },
          {
            path: "members",
            element: <ClusterMembers />
          },
          {
            path: "settings",
            element: <ClusterSettings />
          }
        ]
      },
      {
        path: "pinned",
        element: <Pinned />
      },
      {
        path: "shared",
        element: <Shared />
      },
      {
        path: "settings",
        element: <Settings />
      },
      {
        path: "create",
        element: <CreateLayout />,
        children: [
          {
            element: <OnboardingName />,
            index: true
          },
          {
            path: "onboarding-admins",
            element: <OnboardingAdmins />,
          },
          {
            path: "onboarding-access",
            element: <OnboardingAccess />
          },
          {
            path: "nft-selection",
            element: <NFTSelection />
          },
          {
            path: "address-selection",
            element: <AddressSelection />
          },
          {
            path: "onboarding-invite",
            element: <OnboardingInvite />
          }
        ]
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

  if (address) {
    return (
      <RouterProvider router={mainRouter} />
    )
  }

  return (
    <RouterProvider router={welcomeRouter} />
  );
}

export default App;
