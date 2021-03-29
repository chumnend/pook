import { useEffect, useRef } from 'react';
import { Switch, Route, Link } from 'react-router-dom';

import { useAuth } from './context/auth';

import Home from './pages/Home';
import Landing from './pages/Landing';
import NotFound from './pages/NotFound';
import Register from './pages/Register';
import Login from './pages/Login';
import Logout from './pages/Logout';

import styles from './App.module.css';

const App = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.getToken();
  }, []);

  return (
    <div className={styles.wrapper}>
      <h1>Pook</h1>

      <ul>
        <Link to="/home">Home</Link>
        <Link to="/register">Register</Link>
        <Link to="/login">Login</Link>
        <Link to="/logout">Logout</Link>
        <Link to="/">Landing</Link>
      </ul>

      <Switch>
        <Route exact path="/home" component={Home} />
        <Route exact path="/register" component={Register} />
        <Route exact path="/login" component={Login} />
        <Route exact path="/logout" component={Logout} />
        <Route exact path="/" component={Landing} />
        <Route component={NotFound} />
      </Switch>
    </div>
  );
};

export default App;
