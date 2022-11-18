import TextInput from "../components/TextInput";
import { useState } from "react";

function Home() {
  const [search, setSearch] = useState<string>("");

  function handleSearch() {
    console.log("searching")
  }
  return (
    <div>
      <div className="flex">
        <h1 className="text-4xl font-bold text-onboarding-bg">BanklessDAO</h1>
      </div>
      <div className="bg-white flex px-6 py-8 mt-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        <div className="text-center">
          <h2>Files pinned</h2>
          <p className="mt-2 font-bold text-3xl text-center">1</p>
        </div>
        <div className="text-center">
          <h2>Data used</h2>
          <p className="mt-2 font-bold text-3xl">4.3 MB</p>
        </div>
        <div className="text-center">
          <h2>Total members</h2>
          <p className="mt-2 font-bold text-3xl">14</p>
        </div>
        <div className="text-center">
          <h2>Active members</h2>
          <p className="mt-2 font-bold text-3xl">1</p>
        </div>
      </div>
      <div className="flex mt-6">
        <h2 className="font-bold text-3xl mr-6">Files</h2>
        <TextInput placeholder="Search for file" onChange={handleSearch} />
      </div>
    </div>
  );
}

export default Home;
