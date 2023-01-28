import './App.css';
import Store from './components/Store';
import Header from "./components/header";
import AddCourse from './components/AddCourse';

function App() {
  return (
    <div className="App">
      <Header />
      <AddCourse />
      <Store />
    </div>
  );
}

export default App;
