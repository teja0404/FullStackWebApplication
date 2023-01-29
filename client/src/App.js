import './App.css';
import Store from './components/Store';
import Header from "./components/header";
import AddCourse from './components/AddCourse';
import PaymentHistory from'./components/PaymentHistory';

function App() {
  return (
    <div className="App">
      <Header />
      <AddCourse />
      <PaymentHistory />
      <Store />
    </div>
  );
}

export default App;
