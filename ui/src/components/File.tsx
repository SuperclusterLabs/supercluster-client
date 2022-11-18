import React, { useState } from "react";
import { TrashCanIcon } from "./TrashCanSvg";

export const File = (props: any) => {
  const [contents, setContents] = useState(props.file.contents);
  const name = props.file.name;
  const [showDelete, setShowDelete] = useState(false);

  async function handleSetContents() {
    const payload = {
      contents: contents,
    };

    const resp = await fetch(`/api/files/${props.file.name}`, {
      method: "PUT",
      body: JSON.stringify(payload),
    });

    const body = await resp.json();
    const { file } = body;
    setContents(file.contents);
  }

  async function handleDelete(_: any) {
    await fetch(`/api/files/${props.file.name}`, { method: "DELETE" });
    props.onDeleteSuccess();
  }
  return (
    <div
      className="file"
      onMouseEnter={() => setShowDelete(true)}
      onMouseLeave={() => setShowDelete(false)}
    >
      <p>{name}</p>
      <input
        type={"text"}
        className=""
        onChange={(e) => setContents(e.target.value)}
        value={contents}
      />
      {showDelete && (
        <div>
          <button className="delete" onClick={handleSetContents}>
            Modify Contents
          </button>
          <button className="delete" onClick={handleDelete}>
            <TrashCanIcon />
          </button>
        </div>
      )}
    </div>
  );
};
