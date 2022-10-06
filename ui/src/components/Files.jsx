import React, { useState } from "react";
import { useCallback } from "react";
import { useEffect } from "react";
import { NewFileInput } from "./NewFileForm";
import { File } from "./File";
import { Link } from "react-router-dom";

export const Files = () => {
  const [files, setFiles] = useState([]);

  const fetchFiles = useCallback(async () => {
    const resp = await fetch("/api/files");
    const body = await resp.json();
    const { files } = body;
    console.log(files)
    setFiles(files);
  }, [setFiles]);

  useEffect(() => {
    fetchFiles();
  }, [fetchFiles]);

  function onDeleteSuccess() {
    fetchFiles();
  }

  function onCreateSuccess(newFile) {
    setFiles([...files, newFile]);
  }

  return (
    <>
      <h3>Store:</h3>
      <div className="files">
        {files.map((file) => (
          <File key={file.name} file={file} onDeleteSuccess={onDeleteSuccess} />
        ))}
      </div>
      <NewFileInput onCreateSuccess={onCreateSuccess} />
      <Link to="/about" className="nav-link">
        Learn more...
      </Link>
    </>
  );
};
