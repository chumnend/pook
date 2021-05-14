import PropTypes from 'prop-types';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import ProtectedRoute from '../ProtectedRoute';

export const LANDING_ROUTE = '/';
export const HOME_ROUTE = '/home';
export const REGISTER_ROUTE = '/register';
export const LOGIN_ROUTE = '/login';
export const LOGOUT_ROUTE = '/logout';

const Router = ({ user = null }) => {
  const isSignedIn = user !== null;

  return (
    <BrowserRouter>
      <Switch>
        {/** authenticated routes */}
        <ProtectedRoute
          path={HOME_ROUTE}
          component={() => <h1>Home</h1>}
          condition={isSignedIn}
          redirect={LOGIN_ROUTE}
        />
        <ProtectedRoute
          path={LOGOUT_ROUTE}
          component={() => <h1>Register</h1>}
          condition={isSignedIn}
          redirect={HOME_ROUTE}
        />

        {/** unauthenticated routes */}
        <ProtectedRoute
          path={LOGIN_ROUTE}
          component={() => <h1>Sign In</h1>}
          condition={!isSignedIn}
          redirect={HOME_ROUTE}
        />
        <ProtectedRoute
          path={REGISTER_ROUTE}
          component={() => <h1>Register</h1>}
          condition={!isSignedIn}
          redirect={HOME_ROUTE}
        />

        {/** default routes */}
        <Route exact path={LANDING_ROUTE} component={() => <h1>Landing</h1>} />
        <Route component={() => <h1>Not Found</h1>} />
      </Switch>
    </BrowserRouter>
  );
};

Router.propTypes = {
  user: PropTypes.object,
};

export default Router;
