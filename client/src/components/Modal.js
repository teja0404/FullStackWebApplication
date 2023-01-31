import React,{ useState } from "react";
import "./styles/Header.css";
import axios from "axios";

function Modal({ setOpenModal, cartTotal, finalCourses }) {

  const [name, setName] = useState([]);

  const setModalclose = () => {
    setOpenModal(false);
  }

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Inside Post API for make payment");
    
    e.preventDefault();
    if(cartTotal === ""){
    alert("cartTotal is zero or null");
    }else{
      var options = {
        key: "",
        key_secret:"",
        amount: cartTotal *100,
        currency:"INR",
        name:"Udemy",
        description:"for testing purpose",
        handler: function(response){
          alert(response.razorpay_payment_id);
        },
        prefill: {
          name:"Teja",
          email:"tejailla0404@gmail.com",
          contact:"9110321999"
        },
        notes:{
          address:"Ganesh nagar, Tadepalligudem"
        },
        theme: {
          color:"#3399cc"
        }
      };
      var pay = new window.Razorpay(options);
      pay.open();
    }

    axios
       .post('http://localhost:8081/makepayment', {
        name: name,
        bill : cartTotal,
        courses : finalCourses
       })
       .catch((err) => {
          console.log(err);
       });
 };

 

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
        <div>
        <form onSubmit={handleSubmit}>
        <input
            type="text"
            value={name}
            placeholder="Enter your Name to get invoice on this customer name"
            onChange={(e) => setName(e.target.value)}
          />
          <button type="submit" onClick={() => { setTimeout(setModalclose, 500); } }>
            Pay </button>
        </form>
      </div>
        <div className="footer">
          <button
            onClick={() => {setOpenModal(false);}}
            id="cancelBtn">
            Cancel
          </button>

        </div>

      </div>
    </div>
  );
}

export default Modal;