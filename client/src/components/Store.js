import React, { useState, useEffect } from "react";
import axios from "axios";
import "./styles/List.css";
import { v4 as uuidv4 } from 'uuid';

const Store = () => {
    const [cart, setCart] = useState([]);
    const [items, setItems] = useState([]);
    const [cartTotal, setCartTotal] = useState(0);
    const [payButton, setPayButton] = useState(true)


    useEffect(() => {
        getAllOrders();
    }, [])

    const addToCart = (el) => {
        setCart([...cart, el]);
    };

    useEffect(() => {
        total();
      }, [cart]);
    
      const total = () => {
        let totalVal = 0;
        for (let i = 0; i < cart.length; i++) {
          totalVal += cart[i].price;
        }
        setCartTotal(totalVal);
        setPayButton(totalVal > 0 ? false : true);
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
        <div key={el.id}>
          {`${el.name}: ${el.instructorName}`}
          <input className="list-input" type="submit" value="remove" onClick={() => removeFromCart(el)} />
        </div>
      ));
    
     const listItems = items.map((el) => (
        <div key={el.id} className = "list">
          {`${el.name}: ${el.instructorName} : ${el.price}`}
          <input className="list-input" type="submit" value="add" onClick={() => addToCart(el)} />
        </div>
    ));

    const removeFromCart = (el) => {
        let hardCopy = [...cart];
        hardCopy = hardCopy.filter((cartItem) => cartItem.id !== el.id);
        setCart(hardCopy);
      };

    


      return (
        <div>
          <div><h1>Courses</h1></div>
          <div>{listItems}</div>
          <div><h1>Your cart</h1></div>
          <div>{cartItems}</div>
          <div>Total: {cartTotal}</div>
          <div>
            <button disabled={payButton}>
              Proceed
            </button>
          </div>
        </div>
      );
}

export default Store;