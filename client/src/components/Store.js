import React, { useState, useEffect } from "react";
import axios from "axios";
import "./styles/List.css";
import Modal from "./Modal"
import "./styles/Store.css";

const Store = () => {
    const [cart, setCart] = useState([]);
    const [items, setItems] = useState([]);
    const [cartTotal, setCartTotal] = useState(0);
    const [payButton, setPayButton] = useState(true)
    const [payment, setPayment] = useState(false)
    const [show, setShow] = useState(false);
    const [modalOpen, setModalOpen] = useState(false);
    const [finalCourses, setFinalCourses] = useState();

    var tempString;
    var tempFinalString;

    useEffect(() => {
        getAllOrders();
    }, [])

    const addToCart = (el) => {
      let addToCart = true;

      for(let i=0; i< cart.length; i++) {
        if(cart[i].id == el.id) {
          addToCart = false;
        }
      }

      if(addToCart) {
        setCart([...cart, el]);
      } else {
        alert(`${el.name} by ${el.instructorName} is already in the cart. Please add another course`)
      }
    };

    useEffect(() => {
        total();
      }, [cart]);
    
      const total = () => {
        let totalVal = 0;
        for (let i = 0; i < cart.length; i++) {
          totalVal += cart[i].price;
          tempString = cart[i].name + " by "+ cart[i].instructorName;
          tempFinalString = finalCourses;
          if(tempFinalString == null) {
            tempFinalString = tempString;
          } else {
            tempFinalString = tempFinalString +" , "+ tempString;
          }
        }
        setCartTotal(totalVal);
        setPayButton(totalVal > 0 ? false : true);
        setFinalCourses(tempFinalString);
      };
    
    function getAllOrders(){
        var url = "http://localhost:8081/getallcourses"
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if(response.status == 200){
                setItems(response.data.allCourses)
            }
        })
    }

    const cartItems = cart.map((el) => (
        <div class = "list-items" key={el.id}>
          {`${el.name} by ${el.instructorName}`}
          <input className="list-input" type="submit" value="remove from cart" onClick={() => removeFromCart(el)} />
        </div>
      ));
    
     const listItems = items.map((el) => (
        <div key={el.id} className = "list">
          <div className = "list-details">
              <div className = "list-details-title"><strong>Name</strong> : {`${el.name}`}, </div>
              <div className = "list-details-title"><strong>Instructor Name</strong> : {`${el.instructorName}`}, </div>
              <div className = "list-details-title"><strong>Price</strong> : {`${el.price}`}/-</div>
          </div>
          <div>
          <input className="list-input" type="submit" value="add to cart" onClick={() => addToCart(el)} />
          </div>
        </div>
    ));

    const removeFromCart = (el) => {
        let hardCopy = [...cart];
        hardCopy = hardCopy.filter((cartItem) => cartItem.id !== el.id);
        setCart(hardCopy);
      };

      return (
        <div className="store">
          <div className="Stores-body"><h1>Available Courses</h1></div>
          <div>{listItems}</div>
          <div><h1>Your cart</h1></div>
          <div>{cartItems}</div>
          <div class = "total">Total: {cartTotal}</div>
          <div>
            <button class = "payment" disabled={payButton} onClick={() => {setModalOpen(true)}}>
              Proceed
            </button>
          </div>
          {modalOpen && <Modal  setOpenModal={setModalOpen} cartTotal = {cartTotal} finalCourses = {finalCourses}/>}
        </div>
      );
}

export default Store;