import React from "react";
import { useState } from "react";
import axios from "axios";

function PaymentModal ({setPaymentModalOpen}) {
    const [name, setName] = useState("");
    const [payments, setPayments] = useState([]);
    const [paymentsHistory, setPaymentsHistory] = useState([false]);
    const [showOkayButton, setShowOkayButton] = useState([true]);


    const paymentsList = payments.map((el) => (
        <div key={el.id} className = "list">
          {`Name : ${el.name}, Courses: ${el.courses}, bill: ${el.bill}/-, Date Of Purchase: ${el.Date}`}
        </div>
    ));

    const handleSubmit = (e) => {
          e.preventDefault();
          setShowOkayButton(false)
          axios
             .get('http://localhost:8081/getpaymentsbyname/'+name)
             .then((res) => {
              setName('');
              if(res.status == 200){
                setPayments(res.data.payments)
              }
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
          <div className="body">
            <p>Please enter the name of the customer to get the details</p>
          </div>
      <div>
        <form onSubmit={handleSubmit}>
        <input
            type="text"
            value={name}
            placeholder="Customer Name"
            onChange={(e) => setName(e.target.value)}
          />
  
          <button type="submit">Get Transaction History for this customer</button>
        </form>
      </div>
          <div className="footer">
            {/* <button
              onClick={() => {
                setPaymentModalOpen(false);
                setPaymentsHistory(true);
              }}
              id="cancelBtn"
            >
              Cancel
            </button> */}
            <button  disabled={showOkayButton} onClick={() => {
                  setPaymentModalOpen(false);
                  setPaymentsHistory(true);
              }}>
              Okay
            </button>
          </div>
        </div>
        <div>{paymentsList}</div>
      </div>
    );
}

export default PaymentModal;