import React from "react";
import "./styles/Header.css";
import { useState } from "react";
import axios from "axios";



function AddCourseModal({ setCourseModalOpen }) {
    const [name, setName] = useState("");
    const [instructorName, setInstructorName] = useState("");
    const [price, setPrice] = useState("");
    const [description, setDescription] = useState("");
    const [duration, setDuration] = useState("");

    const setModalclose = () => {
      setCourseModalOpen(false);
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        axios
           .post('http://localhost:8081/addcourse', {
            
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
     };


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
          style = {{margin : "10px", height : "20px", width : "250px"}}
          value={name}
          placeholder="Course Name"
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          style = {{margin : "10px", height : "20px", width : "250px"}}
          value={instructorName}
          placeholder="InstructorName"
          onChange={(e) => setInstructorName(e.target.value)}
        />
         <input
          type="text"
          style = {{margin : "10px", height: "20px", width : "250px"}}
          value={price}
          placeholder="Price"
          onChange={(e) => setPrice(e.target.value)}
        />

        <input
          type="text"
          style = {{margin : "10px", height:"20px", width : "250px"}}
          value={description}
          placeholder="Description"
          onChange={(e) => setDescription(e.target.value)}
        />       
        <input
          type="text"
          style = {{margin : "10px", height : "20px", width : "250px"}}
          value={duration}
          placeholder="Duration"
          onChange={(e) => setDuration(e.target.value)}
        />  
        <button type="submit"  onClick={() => { setTimeout(setModalclose, 500); }}>
          Create</button>
      </form>
    </div>
        <div className="footer">
          <button
            onClick={() => {
                setCourseModalOpen(false);
            }}
            style = {{backgroundColor : "white", width : "200px", color:"black", marginTop: "10px", marginLeft:"100px"}}
            id="cancelBtn"
          >
            Cancel
          </button>
          {/* <button onClick={() => {
                setCourseModalOpen(false);
            }}>
            Ok
          </button> */}
        </div>
      </div>
    </div>
  );
}

export default AddCourseModal;