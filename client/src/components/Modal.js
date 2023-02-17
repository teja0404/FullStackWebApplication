import React, { useCallback, useEffect, useState } from "react";
import "./styles/Header.css";
import CheckoutForm from "./CheckoutForm";
import "./styles/Stripe.css"

import { loadStripe } from "@stripe/stripe-js";
import {
  Elements
} from "@stripe/react-stripe-js";

function Modal({ setOpenModal, cartTotal, finalCourses, name }) {
  const [message, setMessage] = useState("");
  const [receivedSecret, setReceivedSecret] = useState(false);
  const [clientSecret, setClientSecret] = useState("");

  const socket = new WebSocket("wss://fullstackwebapplication-server-nvywvuaozq-uc.a.run.app/getclientsecret");

  let connect = () => {
    console.log("Attempting Connection...");
    socket.onopen = () => {
      console.log("Successfully Connected");
      if (!receivedSecret) {
        sendMsg(
          JSON.stringify({ name: name, bill: cartTotal, courses: finalCourses })
        );
      }
    };

    socket.onmessage = (msg) => {
      console.log("Received the below message from server");
      console.log(msg.data);
      if (msg.data != null) {
        setClientSecret(msg.data);
        setReceivedSecret(true);
      }
    };

    socket.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = (error) => {
      console.log("Socket Error: ", error);
    };
  };

  let sendMsg = (msg) => {
    console.log("sending msg: ", msg);
    socket.send(msg);
  };

  const stripePromise = loadStripe(
    process.env.REACT_APP_STRIPE_PUBLISHABLE_KEY
  );

  useEffect(() => {
    if (!receivedSecret) {
      connect();
    }
  });


  const setModalclose = () => {
    setOpenModal(false);
  };

  const appearance = {
    theme: "stripe",
  };
  const options = {
    clientSecret,
    appearance,
  };

  return (
    <div className="modalBackground">
      <div className="modalContainer">
        <div className="title">
          <h1>Are You Sure You Want to proceed? </h1>
        </div>
        <div className="body">
          <p style = {{color : "whitesmoke"}}>Your final billing amount is {cartTotal} /-</p>
        </div>
        <div className="form">

          <div className="Stripe">
            {clientSecret && (
              <Elements options={options} stripe={stripePromise}>
                <CheckoutForm cartTotal = {cartTotal} finalCourses = {finalCourses} name = {name} />
              </Elements>
            )}
          </div>
          <button
            onClick={() => {
              setOpenModal(false);
            }}
            style={{ width: "100px", marginTop: "5px", marginBottom: "10px" }}
            id="cancelBtn"
          >
            Cancel
          </button>

        </div>
      </div>
    </div>
  );
}

export default Modal;
