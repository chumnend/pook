import { useEffect } from 'react';
import { Switch, Route, Link } from 'react-router-dom';

import Home from './containers/Home';
import Landing from './containers/Landing';
import NotFound from './containers/NotFound';
import Register from './containers/Register';
import Login from './containers/Login';
import Logout from './containers/Logout';

import styles from './App.module.css';

const App = () => {
  useEffect(() => {
    const prefix = process.env.REACT_APP_API_PREFIX;
    const url = prefix + '/health';

    fetch(url)
      .then((res) => res.json())
      .then((data) => console.log(data))
      .catch((err) => console.error(err));
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
