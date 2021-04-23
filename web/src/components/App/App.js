import { useEffect, useRef } from 'react';
import { Route, Switch } from 'react-router-dom';

import * as ROUTES from '../../common/constants/routes';
import useAuth from '../../common/hooks/useAuth';
import Home from '../Home';
import Landing from '../Landing';
import Login from '../Login';
import Logout from '../Logout';
import NotFound from '../NotFound';
import Register from '../Register';
import ProtectedRoute from './components/ProtectedRoute';

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
