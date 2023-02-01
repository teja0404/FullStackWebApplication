import React,{ useState } from "react";
import "./styles/Header.css";
import axios from "axios";
import GooglePayButton from '@google-pay/button-react';

function Modal({ setOpenModal, cartTotal, finalCourses }) {

  const [name, setName] = useState([]);
  let billAmount = cartTotal;
  const [paymentOk, setPaymentOk] = useState(false);
 
  const setModalclose = () => {
    setOpenModal(false);
  }

  const handleSubmit = () => {
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

 const refreshPage = () => { 
  alert("Payment successful")
  window.location.reload(); 
}

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
      <GooglePayButton
        environment="TEST"
        paymentRequest={{
          apiVersion: 2,
          apiVersionMinor: 0,
          allowedPaymentMethods: [
            {
              type: 'CARD',
              parameters: {
                allowedAuthMethods: ['PAN_ONLY', 'CRYPTOGRAM_3DS'],
                allowedCardNetworks: ['MASTERCARD', 'VISA'],
              },
              tokenizationSpecification: {
                type: 'PAYMENT_GATEWAY',
                parameters: {
                  gateway: 'example',
                  gatewayMerchantId: 'exampleGatewayMerchantId',
                },
              },
            },
          ],
          merchantInfo: {
            merchantId: '12345678901234567890',
            merchantName: 'Demo Merchant',
          },
          transactionInfo: {
            totalPriceStatus: 'FINAL',
            totalPriceLabel: 'Total',
            totalPrice: billAmount.toString(),
            currencyCode: 'INR',
            countryCode: 'US',
          },
          shippingAddressRequired: true,
          callbackIntents: ['SHIPPING_ADDRESS', 'PAYMENT_AUTHORIZATION'],
        }}
        onLoadPaymentData={paymentRequest => {
          console.log('Success', paymentRequest);
          setPaymentOk(true);
        }}
        onPaymentAuthorized={paymentData => {
          handleSubmit()
            console.log('Payment Authorised Success', paymentData)
            return { transactionState: 'SUCCESS'}
          }
        }
        onPaymentDataChanged={paymentData => {
            console.log('On Payment Data Changed', paymentData)
            return { }
          }
        }
        existingPaymentMethodRequired='false'
        buttonColor='black'
        buttonType='Buy'
      />
    </div>
  );
}

export default Modal;