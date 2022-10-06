export const Modal = (props) => {
  const showHideClassName = props.show ? "modal display-block" : "modal display-none";

  async function handleClose(e) {
    e.preventDefault();
    props.handleClose();
  }

  return (
    <div className={showHideClassName}>
      <section className="modal-main">
        {props.children}
        <button type="button" onClick={(e) => handleClose(e)}>
          Close
        </button>
      </section>
    </div>
  );
};
