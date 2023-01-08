import { useState } from "react";
import { useNavigate } from "react-router-dom";
import ButtonPrimary from "../components/ButtonPrimary";

function OnboardingAccess() {
  const [accessControl, setAccessControl] = useState<string>("");

  const navigate = useNavigate();

  function confirmAccess() {
    switch (accessControl) {
      case "nft":
        navigate("../nft-selection");
        break;
      case "token":
        navigate("../token-selection");
        break;
      case "addresses":
        navigate("../address-selection");
        break;
    }
  }

  return (
    <div className="text-l-slateblue-700 mt-6">
      <h1 className="text-4xl font-bold">
        Nice! ❤️ Next, let’s set up your cluster’s access controls.
      </h1>
      <p className="text-xl my-4">
        What type of access control does your team use? We’ll use this
        information to make sure your team can get the files they need.
      </p>
      <div className="py-8 columns-3">
        <div
          className={`cursor-pointer rounded-3xl drop-shadow text-center p-8 ${accessControl === "token"
            ? "text-white bg-l-slateblue-700"
            : "text-l-slateblue-700 bg-white"
            }`}
          onClick={() => setAccessControl("token")}
        >
          <h2 className="text-2xl font-bold mb-6">Token Gating</h2>
          <p>Members can access files if they have a certain amount of your DAO's tokens.</p>
        </div>
        <div
          className={`cursor-pointer rounded-3xl drop-shadow text-center p-8 ${accessControl === "nft"
            ? "text-white bg-l-slateblue-700"
            : "text-l-slateblue-700 bg-white"
            }`}
          onClick={() => setAccessControl("nft")}
        >
          <h2 className="text-2xl font-bold mb-6">NFT Gating</h2>
          <p>Members can access files if they own a specific NFT.</p>
        </div>
        <div
          className={`cursor-pointer rounded-3xl drop-shadow text-center p-8 ${accessControl === "addresses"
            ? "text-white bg-l-slateblue-700"
            : "text-l-slateblue-700 bg-white"
            }`}
          onClick={() => setAccessControl("addresses")}
        >
          <h2 className="text-2xl font-bold mb-6">ETH Addresses</h2>
          <p>Provide a list of addresses that can access all your files.</p>
        </div>
      </div>
      <ButtonPrimary onClick={confirmAccess} text="Confirm access" />
    </div>
  );
}

export default OnboardingAccess;
