import React from 'react'

import "./styles/PaymentsSuccess.css"

function PaymentSuccess() {
    return (
        <div className = "success-top">
        <div className="success-body">
        <div className="card-size">
          <i class="checkmark">✓</i>
        </div>
          <h1>Payment Successful</h1> 
          <h1>Thank you for shopping at Udemy</h1>
          <form action="http://localhost:3000/">
            <input type="submit" value="Go to Home" />
            </form>
        </div>
        </div>
        
        
    )
}

export default PaymentSuccess