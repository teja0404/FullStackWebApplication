import "./styles/PaymentsSuccess.css"

import React from "react";

function PaymentSuccess() {
    return (
        <div className = "success-top">
        <div className="success-body">
        <div className="card-size">
          <i class="checkmark">✓</i>
        </div>
          <h1>Payment Successful</h1> 
          <h1>Thank you for shopping at Udemy</h1>
          <form action="https://fullstackwebapplication-client-nvywvuaozq-uc.a.run.app/" style={{paddingRight : "50px"}}>
            <input style = {{margin: "20px", height : "50px", width : "300px", borderRadius : "4px"}} type="submit" value="Go to Home" />
            </form>
        </div>
        </div>
        
        
    )
}

export default PaymentSuccess