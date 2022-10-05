import React, { useState } from "react";

import { Modal } from "./Modal";

export const NewFileInput = (props) => {
  const [showInput, setShowInput] = useState(false);
  const [name, setName] = useState("");
  const [value, setValue] = useState("");
  const [showModal, setShowModal] = useState(false);

  function handleClose() {
    setShowModal(false);
  }
  async function handleSave(e) {
    e.preventDefault();
    if (name.length === 0) {
      setShowModal(true);
      return;
    }

    const payload = {
      name: name,
      contents: value,
    };

    const resp = await fetch("/api/files", {
      method: "POST",
      body: JSON.stringify(payload),
    });

    const body = await resp.json();
    console.log({ body });

    props.onCreateSuccess(body.file);
    setShowInput(false);
    setValue("");
  }

  return (
    <div>
      <Modal show={showModal} handleClose={handleClose}>
        <p>Name cannot be empty</p>
      </Modal>
      {showInput ? (
        <form className="input-box">
          <input
            className="file-input"
            autoFocus
            value={name}
            onChange={(e) => setName(e.target.value)}
            type="text"
            placeholder="file name"
          />
          <input
            className="file-input"
            autoFocus
            value={value}
            onChange={(e) => setValue(e.target.value)}
            type="text"
            placeholder="file contents"
          />
          <button className="save-button" onClick={(e) => handleSave(e)}>
            Save
          </button>
        </form>
      ) : (
        <div className="button-box">
          <button className="new-button" onClick={() => setShowInput(true)}>
            New
          </button>
        </div>
      )}
    </div>
  )
};
