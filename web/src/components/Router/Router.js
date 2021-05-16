import PropTypes from 'prop-types';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import LandingPage from '../LandingPage';
import NotFoundPage from '../NotFoundPage';
import ProtectedRoute from '../ProtectedRoute';
import SignInPage from '../SignInPage';
import SignUpPage from '../SignUpPage';

export const HOME_ROUTE = '/';
export const REGISTER_ROUTE = '/register';
export const LOGIN_ROUTE = '/login';
export const LOGOUT_ROUTE = '/logout';
export const NOT_FOUND_ROUTE = '/not-found';

const Router = ({ user = null }) => {
  const isSignedIn = user !== null;

  return (
    <BrowserRouter>
      <Switch>
        {/** authenticated routes */}
        <ProtectedRoute
          path={LOGOUT_ROUTE}
          component={() => <h1>Register</h1>}
          condition={isSignedIn}
          redirect={HOME_ROUTE}
        />

        {/** unauthenticated routes */}
        <ProtectedRoute
          path={LOGIN_ROUTE}
          component={SignInPage}
          condition={!isSignedIn}
          redirect={HOME_ROUTE}
        />
        <ProtectedRoute
          path={REGISTER_ROUTE}
          component={SignUpPage}
          condition={!isSignedIn}
          redirect={HOME_ROUTE}
        />

        {/** default routes */}
        <Route
          exact
          path={HOME_ROUTE}
          component={isSignedIn ? () => <h1>Home</h1> : LandingPage}
        />
        <Route component={NotFoundPage} />
      </Switch>
    </BrowserRouter>
  );
};

Router.propTypes = {
  user: PropTypes.object,
};

export default Router;
