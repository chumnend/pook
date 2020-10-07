import React, { useEffect, useContext } from 'react';
import { Switch, Route } from 'react-router-dom';

import UserBar from './components/UserBar';
import NavBar from './components/NavBar';
import Footer from './components/Footer';
import Landing from './containers/Landing';
import Register from './containers/Register';
import Login from './containers/Login';
import NotFound from './containers/NotFound';
import { AuthContext } from './context/auth';

function App() {
  const { authState, setUser, logout } = useContext(AuthContext);

  useEffect(() => {
    setUser();
  });

  return (
    <>
      <UserBar isLoggedIn={authState.isLoggedIn} logout={logout} />
      <NavBar />
      <Switch>
        <Route exact path="/register" component={Register} />
        <Route exact path="/login" component={Login} />
        <Route exact path="/" component={Landing} />
        <Route component={NotFound} />
      </Switch>
      <Footer />
    </>
  );
}

export default App;
