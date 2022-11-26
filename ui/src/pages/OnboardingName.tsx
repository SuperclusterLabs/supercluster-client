import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import axios from "axios"
import { useAppStore } from "../store/app"

function OnboardingName() {
  const navigate = useNavigate();
  const address = useAppStore((state) => state.address)
  const setCreatedCluster = useAppStore((state) => state.setCreatedCluster)

  const [clusterName, setClusterName] = useState<string>("");

  async function confirmName() {
    let data = {
      "name": clusterName,
      "creator": address
    }

    const config = {
      method: 'post',
      url: 'http://localhost:3000/api/cluster',
      headers: {
        'Content-Type': 'text/plain'
      },
      data: data
    };

    axios(config)
      .then(function(response) {
        console.log(JSON.stringify(response.data));
        setCreatedCluster(response.data)
        navigate("onboarding-admins");
      })
      .catch(function(error) {
        console.log(error);
      });
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setClusterName(e.target.value);
  }

  return (
    <div className="text-[#334574] mt-6">
      <h1 className="text-2xl font-bold mt-4">
        What should we name your cluster?
      </h1>
      <p className="text-lg my-4">
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
