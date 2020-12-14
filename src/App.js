import React from 'react';
import { Switch, Route } from 'react-router-dom';
import Landing from './containers/Landing';
import Auth from './containers/Auth';
import NotFound from './containers/NotFound';
import Header from './components/Header';
import { useAuthContext } from './context/auth';

const headerLinks = [
  {
    label: 'Register',
    url: '/register',
    requiresAuth: false,
  },
  {
    label: 'Login',
    url: '/login',
    requiresAuth: false,
  },
  {
    label: 'Logout',
    url: '/logout',
    requiresAuth: true,
  },
];

const App = () => {
  const authContext = useAuthContext();

  return (
    <>
      <Header links={headerLinks} isLoggedIn={authContext.isLoggedIn} />
      <Switch>
        <Route exact path="/register" render={(props) => <Auth {...props} />} />
        <Route
          exact
          path="/login"
          component={(props) => <Auth {...props} login />}
        />
        <Route
          exact
          path="/logout"
          render={(props) => <Auth {...props} logout />}
        />
        <Route exact path="/" component={Landing} />
        <Route component={NotFound} />
      </Switch>
    </>
  );
};

export default App;
