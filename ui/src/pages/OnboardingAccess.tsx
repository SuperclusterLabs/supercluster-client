import { useState } from "react";
import { useNavigate } from "react-router-dom";
import ButtonPrimary from "../components/ButtonPrimary";
import superclusterLogo from "../assets/superclusterLogo.svg";

function OnboardingAccess() {
  const [accessControl, setAccessControl] = useState<string>("");

  const navigate = useNavigate();

  function confirmAccess() {
    switch (accessControl) {
      case "nft":
        navigate("/nftselection");
        break;
      case "token":
        navigate("/tokenselection");
        break;
      case "addresses":
        navigate("/addressselection");
        break;
    }
  }

  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <h1 className="text-4xl font-bold text-white mb-10">
          Nice! ❤️ Next, let’s set up your cluster’s access controls.
        </h1>
        <p className="text-2xl text-l-slategray-50">
          What type of access control does your team use? We’ll use this
          information to make sure your team can get the files they need.
        </p>
        <div className="py-8 px-10 mt-10 columns-3">
          <div
            className={`rounded-3xl text-center p-8 ${
              accessControl === "token"
                ? "text-l-slateblue-primary bg-l-slateblue-700"
                : "text-l-slateblue-700 bg-l-slateblue-primary"
            }`}
            onClick={() => setAccessControl("token")}
          >
            <h2 className="text-2xl font-bold mb-6">Token Gating</h2>
            <p>Members can access files if they have X amount of DAO tokens.</p>
          </div>
          <div
            className={`rounded-3xl text-center p-8 ${
              accessControl === "nft"
                ? "text-l-slateblue-primary bg-l-slateblue-700"
                : "text-l-slateblue-700 bg-l-slateblue-primary"
            }`}
            onClick={() => setAccessControl("nft")}
          >
            <h2 className="text-2xl font-bold mb-6">NFT Gating</h2>
            <p>Members can access files if they own a specific NFT.</p>
          </div>
          <div
            className={`rounded-3xl text-center p-8 ${
              accessControl === "addresses"
                ? "text-l-slateblue-primary bg-l-slateblue-700"
                : "text-l-slateblue-700 bg-l-slateblue-primary"
            }`}
            onClick={() => setAccessControl("addresses")}
          >
            <h2 className="text-2xl font-bold mb-6">ETH Addresses</h2>
            <p>Provide a list of addresses that can access all your files.</p>
          </div>
        </div>
        <ButtonPrimary onClick={confirmAccess} text="Confirm access" />
      </div>
    </div>
  );
}

export default OnboardingAccess;
