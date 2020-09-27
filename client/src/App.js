import React from 'react';
import { Switch, Route } from 'react-router-dom';

import UserBar from './components/UserBar';
import NavBar from './components/NavBar';
import Footer from './components/Footer';
import Landing from './pages/Landing';
import Register from './pages/Register';
import Login from './pages/Login';
import NotFound from './pages/NotFound';

function App() {
  return (
    <>
      <UserBar />
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
