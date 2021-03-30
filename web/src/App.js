import { useEffect, useRef } from 'react';
import { Switch, Route, Link } from 'react-router-dom';
import { useAuth } from './context/auth';
import ProtectedRoute from './components/ProtectedRoute';
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
        {auth.user && <Link to="/logout">Logout</Link>}

        {auth.user == null && <Link to="/register">Register</Link>}
        {auth.user == null && <Link to="/login">Login</Link>}
      </ul>

      <Switch>
        <ProtectedRoute
          exact
          path="/home"
          component={Home}
          condition={auth.user != null}
          redirect="/login"
        />
        <ProtectedRoute
          exact
          path="/logout"
          component={Logout}
          condition={auth.user != null}
          redirect="/login"
        />
        <ProtectedRoute
          exact
          path="/register"
          component={Register}
          condition={auth.user == null}
          redirect="/home"
        />
        <ProtectedRoute
          exact
          path="/login"
          component={Login}
          condition={auth.user == null}
          redirect="/home"
        />
        <ProtectedRoute
          exact
          path="/"
          component={Landing}
          condition={auth.user == null}
          redirect="/home"
        />
        <Route component={NotFound} />
      </Switch>
    </div>
  );
};

export default App;
