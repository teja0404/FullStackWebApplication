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
          value={name}
          placeholder="Course Name"
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          value={instructorName}
          placeholder="InstructorName"
          onChange={(e) => setInstructorName(e.target.value)}
        />
         <input
          type="text"
          value={price}
          placeholder="Price"
          onChange={(e) => setPrice(e.target.value)}
        />

        <input
          type="text"
          value={description}
          placeholder="Description"
          onChange={(e) => setDescription(e.target.value)}
        />       
        <input
          type="text"
          value={duration}
          placeholder="Duration"
          onChange={(e) => setDuration(e.target.value)}
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