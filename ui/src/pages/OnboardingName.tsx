import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import superclusterLogo from "../assets/superclusterLogo.svg";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";

function OnboardingName() {
  const navigate = useNavigate();

  const [clusterName, setClusterName] = useState<String>("");

  function confirmName() {
    console.log(clusterName);
    navigate("/onboardingadmins");
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setClusterName(e.target.value);
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
          Hey 👋🏼, kaihuang.eth! What should we name your cluster?
        </h1>
        <p className="text-2xl text-l-slategray-50">
          You'll need a name for your cluster. It will help your teammates find
          you a little easier. You can always change this afterwards.
        </p>
        <TextInput onChange={handleInputChange} placeholder="Cluster name" />
        <ButtonPrimary onClick={confirmName} text="Confirm name" />
      </div>
    </div>
  );
}

export default OnboardingName;
