import React from "react";
import "./styles/Header.css";

function Modal({ setOpenModal, cartTotal }) {
  return (
    <div className="modalBackground">
      <div className="modalContainer">
        <div className="titleCloseBtn">
        </div>
        <div className="title">
          <h1>Are You Sure You Want to proceed? </h1>
        </div>
        <div className="body">
          <p>Your final billing amount is {cartTotal}</p>
        </div>
        <div className="footer">
          <button
            onClick={() => {
              setOpenModal(false);
            }}
            id="cancelBtn"
          >
            Cancel
          </button>
          <button>Pay</button>
        </div>
      </div>
    </div>
  );
}

export default Modal;