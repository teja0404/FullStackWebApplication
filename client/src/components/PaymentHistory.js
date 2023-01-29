import React from "react";
import { useState } from "react";
import PaymentModal from "./PaymentModal";
import "./styles/PaymentHistory.css";

import axios from "axios";

const PaymentHistory =  () => {
    const [paymentModalOpen, setPaymentModalOpen] = useState(false);

    return (
    <div>
        <button class = "Payment-History" disabled={paymentModalOpen} onClick={() => {setPaymentModalOpen(true)}}>
          Get Purchase history of customer
        </button>
        {paymentModalOpen && <PaymentModal  setPaymentModalOpen={setPaymentModalOpen}/>}
    </div>
    );
};

export default PaymentHistory;
