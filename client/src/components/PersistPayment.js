import React from "react";
import axios from "axios";

function PersistPayment() {
    e.preventDefault();
    axios
       .post('https://fullstackwebapplication-server-nvywvuaozq-uc.a.run.app/payment', {
        
        name: name,
        instructorName: instructorName,
        price: Number(price),
        description: description,
        duration: Number(duration),
       })
       .then((res) => {
        setName('');
        setInstructorName('');
        setPrice('')
        setDescription('')
        setDuration('')
       })
       .catch((err) => {
          console.log(err);
       });
}

export default PersistPayment;