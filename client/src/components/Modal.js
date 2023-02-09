import React, { useCallback, useEffect, useState } from 'react';
import "./styles/Header.css";
import axios from "axios";
import { connect, sendMsg } from "../websockets/websocket";

function Modal({ setOpenModal, cartTotal, finalCourses}) {

  const [name, setName] = useState([]);
  const [message, setMessage] = useState('')
  // const [ws, setWs] = useState(new WebSocket(URL));
  // const socket = new WebSocket("ws://localhost:8080/makepayment");
  
  let billAmount = cartTotal;

  // const handlePay = () => {
  //   const message = { user: "Teja", message: "Teja is a good boy" };
  //   ws.send(JSON.stringify(message));
  // }



  useEffect(() => {
    connect();
  });

  // const send = () => {
  //   console.log("hello");
  //   sendMsg("hello");
  // }

  function handleSubmit(e) {
    e.preventDefault();
    console.log('You clicked submit.');
    sendMsg("hello");
  }

  // useEffect(() => {
  //   console.log("Hello from function")

  //   ws.onopen = () => {
  //     console.log('WebSocket Connected');
  //   }

  //   ws.onmessage = (e) => {
  //     //Implement receiving messages section
  //   }

  //   return () => {
  //     // ws.onclose = () => {
  //     //   console.log('WebSocket Disconnected');
  //     //   setWs(new WebSocket(URL));
  //     // }
  //   }
  // }, [ws.onmessage, ws.onopen, ws.onclose, []]);
  

 
  const setModalclose = () => {
    setOpenModal(false);
  }


  return (
    <div className="modalBackground">
      <div className="modalContainer">
        <div className="title">
          <h1>Are You Sure You Want to proceed? </h1>
        </div>
        <div className="body">
          <p>Your final billing amount is {cartTotal} /-</p>
        </div>
        <div className="form">
        <form onSubmit={handleSubmit}>
        <input
            type="text"
            value={name}
            style = {{width:"300px"}}
            placeholder="Enter your Name"
            onChange={(e) => setName(e.target.value)}
          />
        </form>
        <button
            onClick={() => {setOpenModal(false);}}
            style = {{width:"100px", marginTop:"5px", marginBottom:"10px"}}
            id="cancelBtn">
            Cancel
        </button>
        <button
            onClick={handleSubmit}
            style = {{width:"100px", marginTop:"5px", marginBottom:"10px"}}
            id="payBtn">
            Pay
        </button>
      </div>
      </div>
    </div>
  );
}

export default Modal;