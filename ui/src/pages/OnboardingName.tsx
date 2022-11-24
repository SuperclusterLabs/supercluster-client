import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";

function OnboardingName() {
  const navigate = useNavigate();

  const [clusterName, setClusterName] = useState<string>("");

  function confirmName() {
    console.log(clusterName);
    navigate("onboarding-admins");
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setClusterName(e.target.value);
  }

  return (
    <div className="text-l-slateblue-700">
      <h1 className="text-2xl font-bold">
        What should we name your cluster?
      </h1>
      <p className="text-l">
        You'll need a name for your cluster. It will help your teammates
        find you a little easier. You can always change this afterwards.
      </p>
      <TextInput
        value={clusterName}
        onChange={handleInputChange}
        placeholder="Cluster name"
      />
      <ButtonPrimary onClick={confirmName} text="Confirm name" />
    </div>
  );
}

export default OnboardingName;
