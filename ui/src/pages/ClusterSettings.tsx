import { useEffect, useState } from "react";

const exampleCluster: any = {
  id: "id_1",
  name: "BanklessDAO",
  description: "File cluster for BanklessDAO Dev Guild",
  nftAddr: "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D",
  files: [{ name: "file" }],
  admins: [{ name: "admin" }],
  creator: "0x6eD68a1982ac2266ceB9C1907B629649aAd9AC20",
  members: [{ name: "member" }]
}

function ClusterSettings() {
  const [cluster, setCluster] = useState<any>(null);

  // TODO: Get cluster data from cluster API endpoint
  useEffect(() => {
    setCluster(exampleCluster);
  }, [])

  if (cluster) {
    return (
      <div className="flex flex-col mt-6 text-l-slateblue-700">
        <div className="mb-5">
          <h2 className="font-bold text-xl mb-2">Cluster name</h2>
          <p>{cluster.name}</p>
        </div>
        <div className="mb-5">
          <h2 className="font-bold text-xl mb-2">Cluster description</h2>
          <p>{cluster.description}</p>
        </div>
        <div className="mb-5">
          <h2 className="font-bold text-xl mb-2">Cluster access</h2>
          <p>{cluster.nftAddr}</p>
        </div>
      </div>
    )
  } else {
    return null
  }
}

export default ClusterSettings;
