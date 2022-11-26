import React, { useState } from "react";
import { Delete } from "react-feather";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";
import { useAppStore } from "../store/app"
import axios from "axios"

function OnboardingAdmins() {
  const [adminList, setAdminList] = useState<Array<string>>([]);
  const [adminAddress, setAdminAddress] = useState<string>("");
  const createdCluster = useAppStore((state) => state.createdCluster)

  const navigate = useNavigate();

  async function confirmAdmins() {
    console.log(createdCluster);
    if (createdCluster) {

      let data = createdCluster;
      data.admins = adminList;

      console.log(data);
      let config = {
        method: 'put',
        url: `http://localhost:3000/api/cluster/${createdCluster.id}`,
        headers: {
          'Content-Type': 'text/plain'
        },
        data: data
      };

      await axios(config)
        .then(function(response: any) {
          console.log(JSON.stringify(response.data));
          navigate("../onboarding-access");
        })
        .catch(function(error: any) {
          console.log(error);
        });
    }

  }

  function addAddress() {
    if (adminList.includes(adminAddress)) {
      return;
    } else {
      setAdminList([...adminList, adminAddress]);
      setAdminAddress("");
    }
  }

  function removeAddress(address: string) {
    setAdminList((current) =>
      current.filter((admin) => {
        return admin !== address;
      })
    );
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    setAdminAddress(e.target.value);
  }

  return (
    <div className="text-[#334574] mt-6">
      <h1 className="text-2xl font-bold mt-4"> ⭐️ Awesome! Who are your cluster’s admins?
      </h1>
      <p className="text-lg my-4">
        Admins can adjust permissions, remove team members, and change your
        cluster’s settings. Make sure you trust them!
      </p>
      <TextInput
        // TODO: Add input validation for ETH addresses
        onChange={handleInputChange}
        placeholder="Enter address or ENS"
        value={adminAddress}
      />
      <ButtonPrimary onClick={addAddress} text="Add admin" />
      <h2 className="text-2xl font-bold mt-8 mb-4">Admin list</h2>
      <div className="container bg-white px-6 py-8 mb-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        {adminList.map((adminAddress: string) => (
          <div className="text-xl mb-2" key={adminAddress}>
            {adminAddress}
            <Delete
              className="ml-2 inline cursor-pointer"
              size={16}
              onClick={() => removeAddress(adminAddress)}
            />
          </div>
        ))}
      </div>
      <ButtonPrimary onClick={confirmAdmins} text="Confirm admins" />
    </div>
  );
}

export default OnboardingAdmins;
