import { useEffect, useRef } from 'react';
import { Route, Switch } from 'react-router-dom';

import ProtectedRoute from './components/ProtectedRoute';
import * as ROUTES from './constants/routes';
import Home from './containers/Home';
import Landing from './containers/Landing';
import Login from './containers/Login';
import Logout from './containers/Logout';
import NotFound from './containers/NotFound';
import Register from './containers/Register';
import { useAuth } from './context/auth';

const App = () => {
  const auth = useAuth();
  const authRef = useRef(auth);

  useEffect(() => {
    authRef.current.getToken();
  }, []);

  return (
    <Switch>
      <ProtectedRoute
        exact
        path={ROUTES.HOME}
        component={Home}
        condition={auth.user != null}
        redirect={ROUTES.LOGIN}
      />
      <ProtectedRoute
        exact
        path={ROUTES.LOGOUT}
        component={Logout}
        condition={auth.user != null}
        redirect={ROUTES.LOGIN}
      />
      <ProtectedRoute
        exact
        path={ROUTES.REGISTER}
        component={Register}
        condition={auth.user == null}
        redirect={ROUTES.HOME}
      />
      <ProtectedRoute
        exact
        path={ROUTES.LOGIN}
        component={Login}
        condition={auth.user == null}
        redirect={ROUTES.HOME}
      />
      <ProtectedRoute
        exact
        path={ROUTES.LANDING}
        component={Landing}
        condition={auth.user == null}
        redirect={ROUTES.HOME}
      />
      <Route component={NotFound} />
    </Switch>
  );
};

export default App;
