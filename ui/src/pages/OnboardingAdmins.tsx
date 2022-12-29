import React, { useState } from "react";
import { Delete } from "react-feather";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";
import { useAppStore } from "../store/app"
import { checkIfPathIsEns, checkIfPathIsEth } from "../helpers/string"
import axios from "axios"

function OnboardingAdmins() {
  const [adminList, setAdminList] = useState<Array<string>>([]);
  const [adminAddress, setAdminAddress] = useState<string>("");
  const [addressErr, setAddressErr] = useState<string>("");
  const createdCluster = useAppStore((state) => state.createdCluster)
  const setCreatedCluster = useAppStore((state) => state.setCreatedCluster)

  const navigate = useNavigate();

  async function confirmAdmins() {
    if (createdCluster) {
      let data = createdCluster;
      data.admins = adminList;

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
          setCreatedCluster(response.data)
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
      if (checkIfPathIsEth(adminAddress) || checkIfPathIsEns(adminAddress)) {
        setAdminList([...adminList, adminAddress]);
        setAdminAddress("");
      } else {
        setAddressErr("Address is not a valid ETH or ENS address")
      }
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

  function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      addAddress()
    }
  }

  return (
    <div className="text-l-slateblue-700 mt-6">
      <h1 className="text-4xl font-bold"> ⭐️ Awesome! Who are your cluster’s admins?
      </h1>
      <p className="text-xl my-4">
        Admins can adjust permissions, remove team members, and change your
        cluster’s settings. Make sure you trust them!
      </p>
      <TextInput
        onChange={handleInputChange}
        onKeyDown={handleKeyDown}
        placeholder="Enter address or ENS"
        value={adminAddress}
        error={addressErr}
      />
      <div className="container bg-white px-6 py-8 my-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        <h2 className="text-2xl font-bold mb-4">Admin list</h2>
        {adminList.map((adminAddress: string) => (
          <div className="text-xl mb-2 table-item" key={adminAddress}>
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
