import React, { useCallback, useEffect, useState } from 'react';
import "./styles/Header.css";
// import "./styles/Stripe.css"
import CheckoutForm from "./CheckoutForm";
//import { connect, sendMsg } from "../websockets/websocket";

import { loadStripe } from "@stripe/stripe-js";
import {
  CardElement,
  Elements,
  useStripe,
  useElements,
} from "@stripe/react-stripe-js";

function Modal({ setOpenModal, cartTotal, finalCourses, name}) {

  // const [name, setName] = useState([]);
  const [message, setMessage] = useState('');
  const [receivedSecret, setReceivedSecret] = useState(false);

  const socket = new WebSocket('ws://localhost:8081/makepayment');

  let connect = () => {
    console.log("Attempting Connection...");
  
    socket.onopen = () => {
      console.log("Successfully Connected");

      if(!receivedSecret) {
        sendMsg(JSON.stringify({ name: name, bill : cartTotal, courses : finalCourses}));
      }
      
    };
  
    socket.onmessage = msg => {
      console.log("Received the below message from server")
      console.log(msg.data);
      if (msg.data != null) {
        setClientSecret(msg.data);
        setReceivedSecret(true);
      }
      
    };
  
    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };
  
    socket.onerror = error => {
      console.log("Socket Error: ", error);
    };
  };

  let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
  };
  
  let billAmount = cartTotal;

  const delay = ms => new Promise(
    resolve => setTimeout(resolve, ms)
  );

  const stripePromise = loadStripe("pk_test_51MZPhbSIDNJpH3u5GphZ07wAcXWpnmHSvDIP2h6SexLh2764zLMke5gv2h158XYwzKanbeeq0zlbIIzpRIUYp6qs00xS822yCO");

  useEffect(() => {
    if(!receivedSecret) {
      connect();
    }
    
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

  // const CheckoutForm = () => {
  //   const stripe = useStripe();
  //   const elements = useElements();
  //   const [error, setError] = useState(null);
  //   const [paymentSuccess, setPaymentSuccess] = useState(false);
  
  //   const handleSubmit = async (event) => {
  //     event.preventDefault();
  
  //     if (!stripe || !elements) {
  //       // Stripe.js has not loaded yet. Make sure to disable
  //       // form submission until Stripe.js has loaded.
  //       return;
  //     }

  //     const { error, paymentMethod } = await stripe.createPaymentMethod({
  //       type: "card",
  //       card: elements.getElement(CardElement),
  //     });
  //     if (!error) {
  //       console.log(paymentMethod);
        
  //       sendMsg(JSON.stringify({ name: name, bill : cartTotal, courses : finalCourses}));

  //       Send the paymentMethod to your server for processing
  //       const socket = new WebSocket("ws://localhost:8081/makepayment");
  //       socket.onopen = () => {
  //         sendMsg(JSON.stringify({ paymentMethod }));
  //         socket.send(JSON.stringify({ paymentMethod }));
  //       };
  //     } else {
  //       setError(error.message);
  //     }
  //   };
  
  //   return (
  //     <form onSubmit={handleSubmit}>
  //       <h2>Payment Information</h2>
  //       <CardElement />
  //       <button type="submit" disabled={!stripe}>
  //         Pay
  //       </button>
  //       {error && <div>{error}</div>}
  //       {paymentSuccess && <div>Payment successful!</div>}
  //     </form>
  //   );
  // };

  const setModalclose = () => {
    setOpenModal(false);
  }

  const [clientSecret, setClientSecret] = useState("");

  useEffect(() => {
    //pi_3MZywrSIDNJpH3u50iVk3Dbn_secret_wUqRsYRkixWBQMyzRo2evFOHY
    // Create PaymentIntent as soon as the page loads
    // fetch("/create-payment-intent", {
    //   method: "POST",
    //   headers: { "Content-Type": "application/json" },
    //   body: JSON.stringify({ items: [{ id: "xl-tshirt" }] }),
    // })
    //   .then((res) => res.json())
    //   .then((data) => setClientSecret(data.clientSecret));


    // if(socket.onopen != null) {
    //   sendMsg(JSON.stringify({ name: name, bill : cartTotal, courses : finalCourses}));
    // }

  }, []);

  const appearance = {
    theme: 'stripe',
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
          <p>Your final billing amount is {cartTotal} /-</p>
        </div>
        <div className="form">
        <button
            onClick={() => {setOpenModal(false);}}
            style = {{width:"100px", marginTop:"5px", marginBottom:"10px"}}
            id="cancelBtn">
            Cancel
        </button>
        <div className="Stripe">
          {clientSecret && (<Elements options={options} stripe={stripePromise}>
          <CheckoutForm />
          </Elements>
           )}
        </div>
        {/* <Elements stripe={loadStripe("pk_test_51MZPhbSIDNJpH3u5GphZ07wAcXWpnmHSvDIP2h6SexLh2764zLMke5gv2h158XYwzKanbeeq0zlbIIzpRIUYp6qs00xS822yCO")}>
          <CheckoutForm />
        </Elements> */}
        {/* <button
            onClick={handleSubmit}
            style = {{width:"100px", marginTop:"5px", marginBottom:"10px"}}
            id="payBtn">
            Pay
        </button> */}
      </div>
      </div>
    </div>
  );
}

export default Modal;