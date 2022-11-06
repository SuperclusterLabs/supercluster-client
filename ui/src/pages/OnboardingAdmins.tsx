import React, { useState } from "react";
import superclusterLogo from "../assets/superclusterLogo.svg";
import { Delete } from "react-feather";
import TextInput from "../components/TextInput";
import ButtonPrimary from "../components/ButtonPrimary";
import { useNavigate } from "react-router-dom";

function OnboardingAdmins() {
  const [adminList, setAdminList] = useState<Array<string>>([]);
  const [adminAddress, setAdminAddress] = useState<string>("");

  const navigate = useNavigate();

  function confirmAdmins() {
    navigate("/onboardingaccess");
  }

  function addAddress() {
    if (adminList.includes(adminAddress)) {
      return;
    } else {
      setAdminList([...adminList, adminAddress]);
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
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <h1 className="text-4xl font-bold text-white mb-10">
          ⭐️ Awesome! Who are your cluster’s admins?
        </h1>
        <p className="text-2xl text-l-slategray-50">
          Admins can adjust permissions, remove team members, and change your
          cluster’s settings. Make sure you trust them!
        </p>
        <TextInput
          onChange={handleInputChange}
          placeholder="Enter address or ENS"
        />
        <ButtonPrimary onClick={addAddress} text="Add admin" />
        <div className="py-8 px-10 bg-l-slateblue-primary w-2/5 rounded-2xl text-left text-l-slateblue-700 mt-10">
          <h2 className="text-2xl font-bold mb-6">Admin list</h2>
          {adminList.map((adminAddress: string) => (
            <div key={adminAddress}>
              <span>
                {adminAddress}{" "}
                <Delete
                  className="inline"
                  size={18}
                  onClick={() => removeAddress(adminAddress)}
                />
              </span>
            </div>
          ))}
          <ButtonPrimary onClick={confirmAdmins} text="Confirm admins" />
        </div>
      </div>
    </div>
  );
}

export default OnboardingAdmins;
