import './App.css';
import Home from './components/Home';
import Login from './components/Login';
import Signup from './components/Signup';

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      {/* <Home /> */}
      <Router>
        <Routes>
            <Route
                exact
                path="/"
                element={<Home />}
            ></Route>
            <Route
                exact
                path="/signup"
                element={<Signup />}
            ></Route>
            <Route
                exact
                path="/login"
                element={<Login />}
            ></Route>
        </Routes>
      </Router>
    </div>
  );
}

export default App;
