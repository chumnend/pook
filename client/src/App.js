import React, { useEffect, useContext, useRef } from 'react';
import { Switch, Route } from 'react-router-dom';

import Header from './components/Header';
import Footer from './components/Footer';
import Landing from './containers/Landing';
import Register from './containers/Register';
import Login from './containers/Login';
import NotFound from './containers/NotFound';
import { AuthContext } from './context/auth';

function App() {
  const auth = useContext(AuthContext);
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.setUser();
  }, []);

  return (
    <>
      <Header isLoggedIn={auth.isLoggedIn} logout={auth.logout} />
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
