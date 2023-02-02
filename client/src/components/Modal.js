import React,{ useState } from "react";
import "./styles/Header.css";
import axios from "axios";
import GooglePayButton from '@google-pay/button-react';

function Modal({ setOpenModal, cartTotal, finalCourses}) {

  const [name, setName] = useState([]);
  let billAmount = cartTotal;
  
 
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