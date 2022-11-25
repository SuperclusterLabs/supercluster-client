import { useEffect, useState } from "react";
import Dropzone from "../components/Dropzone";
import { useAppStore } from "../store/app"

const exampleFiles: Array<any> = [
  {
    name: "Example file 1",
    cid: "zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA",
    type: "mime/image",
    description: "Example description for example file 1",
    size: "43000",
    dateUploaded: "2011-08-12T20:17:46.384Z",
    lastModified: "2021-08-12T20:17:46.384Z",
    uploadedBy: "kaihuang.eth"
  },
  {
    name: "Example file 2",
    cid: "zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA",
    type: "mime/image",
    description: "Example description for example file 2",
    size: "43000",
    dateUploaded: "2011-08-12T20:17:46.384Z",
    lastModified: "2021-08-12T20:17:46.384Z",
    uploadedBy: "kaihuang.eth"
  },
  {
    name: "Example file 3",
    cid: "zb2rhe5P4gXftAwvA4eXQ5HJwsER2owDyS9sKaQRRVQPn93bA",
    type: "mime/image",
    description: "Example description for example file 3",
    size: "43000",
    dateUploaded: "2011-08-12T20:17:46.384Z",
    lastModified: "2021-08-12T20:17:46.384Z",
    uploadedBy: "kaihuang.eth"
  },
]


function ClusterFiles() {
  const cluster = useAppStore((state) => state.activeCluster)
  // TODO: Need to get the files from the Cluster
  const [files, setFiles] = useState<Array<any>>([]);
  const [numberOfFiles, setNumberOfFiles] = useState<number>(0);

  // TODO: Change to get Files from API
  useEffect(() => {
    setFiles(exampleFiles)
    if (cluster) {
      if (cluster.files) {
        setNumberOfFiles(cluster.files.length)
      }

      setFiles(cluster.files)
    }
  }, [])

  function handleFileUpload(e: React.ChangeEvent<HTMLInputElement>) {
    console.log(e.target.files)
  }

  return (
    <div>
      <div className="bg-white flex px-6 py-8 mt-8 rounded-2xl space-x-10 text-l-slateblue-700 drop-shadow">
        <div className="text-center">
          <h2>Files pinned</h2>
          <p className="mt-2 font-bold text-3xl text-center">{numberOfFiles}</p>
        </div>
        <div className="text-center">
          <h2>Total members</h2>
          <p className="mt-2 font-bold text-3xl">14</p>
        </div>
      </div>
      <div className="flex items-center mt-6">
        <h2 className="font-bold text-3xl mr-6">Files</h2>
      </div>
      <div className="flex mt-4">
        <Dropzone multiple={true} onChange={handleFileUpload} />
      </div>
      <table className="table-auto border-separate border-spacing-8 text-l-slateblue-700">
        <thead>
          <tr>
            <th>File type</th>
            <th>Filename</th>
            <th>Size</th>
            <th>Date uploaded</th>
            <th>Last modified</th>
            <th>Uploaded by</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {files && files.map((file, i) => {
            return (
              <tr key={i}>
                <td>{file.type}</td>
                <td>{file.name}</td>
                <td>{file.size}</td>
                <td>{file.dateUploaded}</td>
                <td>{file.lastModified}</td>
                <td>{file.uploadedBy}</td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </div>
  );
}

export default ClusterFiles;

