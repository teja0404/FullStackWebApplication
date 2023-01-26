import React, { useState, useEffect } from "react";
import axios from "axios";
import { v4 as uuidv4 } from 'uuid';

const Store = () => {
    const [cart, setCart] = useState([]);
    const [items, setItems] = useState([]);
    const [cartTotal, setCartTotal] = useState(0);

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
          totalVal += cart[i].Age;
        }
        setCartTotal(totalVal);
      };
    
    function getAllOrders(){
        var url = "http://localhost:8081/getall"
        axios.get(url, {
            responseType: 'json'
        }).then(response => {
            if(response.status == 200){
                setItems(response.data.student)
            }
        })
    }

    const cartItems = cart.map((el) => (
        <div key={el.Name}>
          {`${el.Name}: ${el.Age}`}
          <input type="submit" value="remove" onClick={() => removeFromCart(el)} />
        </div>
      ));
    
     const listItems = items.map((el) => (
        <div key={uuidv4()}>
          {`${el.Name}: ${el.Age}`}
          <input type="submit" value="add" onClick={() => addToCart(el)} />
        </div>
    ));

    const removeFromCart = (el) => {
        let hardCopy = [...cart];
        hardCopy = hardCopy.filter((cartItem) => cartItem.Name !== el.Name);
        setCart(hardCopy);
      };

      return (
        <div>
          udemy
          <div><h1>Courses</h1></div>
          <div>{listItems}</div>
          <div><h1>Your cart</h1></div>
          <div>{cartItems}</div>
          <div>Total: {cartTotal}</div>
        </div>
      );
}

export default Store;