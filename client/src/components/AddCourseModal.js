import React from "react";
import "./styles/Header.css";
import { useState, useEffect } from "react";
import axios from "axios";

function AddCourseModal({ setCourseModalOpen }) {
    const [name, setName] = useState("");
    const [instructorName, setInstructorName] = useState("");
    const [price, setPrice] = useState("");
    const [description, setDescription] = useState("");
    const [duration, setDuration] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        axios
           .post('https://localhost/8081/addcourse', {
            name: "Block chain",
            instructorName: "Suvadip",
            price: 1200,
            description: "Block chain development",
            duration: 120,
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
     };

    // const handleSubmit = (e) => {
    //     e.preventDefault();
    //     console.log("Form submitted");
    //     fetch('https://localhost/8081/addcourse', {
    //        method: 'POST',
    //        body: JSON.stringify({
    //         name: "Block chain",
    //         instructorName: "Suvadip",
    //         price: 1200,
    //         description: "Block chain development",
    //         duration: 120,
    //        }),
    //        headers: {
    //           'Content-type': 'application/json; charset=UTF-8',
    //        },
    //     })
    //  };

    // let handleSubmit = async (e) => {
    //     console.log("Form submitted");
    //     e.preventDefault();
    //     try {
    //       let res = await fetch("https://localhost/8081/addcourse", {
    //         method: "POST",
    //         body: JSON.stringify({
    //           name: "Block chain",
    //           instructorName: "Suvadip",
    //           price: 1200,
    //           description: "Block chain development",
    //           duration: 120,
    //         }),
    //       });
    //       let resJson = await res.json();
    //       if (res.status === 200) {
    //         setName("");
    //         setInstructorName("");
    //         setPrice("")
    //         setDescription("")
    //         setDuration("")

    //       } else {
    //         console.log("Some error while sending Post request");
    //       }
    //     } catch (err) {
    //       console.log(err);
    //     }
    //   };


  return (
    <div className="modalBackground">
      <div className="modalContainer">
        <div className="titleCloseBtn">
        </div>
        <div className="body">
          <p>Please add new course details below </p>
        </div>
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={name}
          placeholder="Name"
          onChange={(e) => setName(e.target.name)}
        />
        <input
          type="text"
          value={instructorName}
          placeholder="InstructorName"
          onChange={(e) => setInstructorName(e.target.instructorName)}
        />
         <input
          type="text"
          value={price}
          placeholder="Price"
          onChange={(e) => setPrice(e.target.price)}
        />
        <input
          type="text"
          value={description}
          placeholder="Description"
          onChange={(e) => setDescription(e.target.description)}
        />       
        <input
          type="text"
          value={duration}
          placeholder="Duration"
          onChange={(e) => setDuration(e.target.duration)}
        />  
        <button type="submit">Create</button>
      </form>
    </div>
        <div className="footer">
          <button
            onClick={() => {
                setCourseModalOpen(false);
            }}
            id="cancelBtn"
          >
            Cancel
          </button>
          <button onClick={() => {
                setCourseModalOpen(false);
            }}>
            Ok
          </button>
        </div>
      </div>
    </div>
  );
}

export default AddCourseModal;