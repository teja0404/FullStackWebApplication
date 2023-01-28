import React from "react";
import "./styles/AddCourse.css";
import AddCourseModal from "./AddCourseModal"
import { useState } from "react";

const AddCourse = () => {
  const [courseModalOpen, setCourseModalOpen] = useState(false);
    return (
        <div>
        <button className="Add-Course" onClick={() => {setCourseModalOpen(true)}}>
          Add course
        </button>
        {courseModalOpen && <AddCourseModal  setCourseModalOpen={setCourseModalOpen} />}
      </div>
    )
}

export default AddCourse;