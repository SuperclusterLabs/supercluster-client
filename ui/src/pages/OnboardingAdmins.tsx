import React, { useState } from "react";
import { Delete } from "react-feather";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";

function OnboardingAdmins() {
  const [adminList, setAdminList] = useState<Array<string>>([]);
  const [adminAddress, setAdminAddress] = useState<string>("");

  const navigate = useNavigate();

  function confirmAdmins() {
    navigate("../onboarding-access");
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
    <div className="text-l-slateblue-700">
      <h1 className="text-2xl font-bold"> ⭐️ Awesome! Who are your cluster’s admins?
      </h1>
      <p className="text-lg">
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
