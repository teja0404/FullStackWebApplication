import './App.css';
import Store from './components/Store';
import Header from "./components/header";
import AddCourse from './components/AddCourse';
import PaymentHistory from'./components/PaymentHistory';
import PaymentSuccess from './components/PaymentSuccess'
import { Routes, Route } from "react-router-dom"

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={ [<Header/>,<AddCourse/>,<PaymentHistory/>,<Store/>]} />
        <Route path="paymentsuccess" element={ <PaymentSuccess/> } />
      </Routes>
    </div>
  );
}

export default App;