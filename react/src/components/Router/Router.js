import React, { Suspense } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import { useAuth } from '../../providers/AuthProvider';
import Loader from '../Loader';
import ProtectedRoute from '../ProtectedRoute';

const HomePage = React.lazy(() => import('../../pages/HomePage'));
const LandingPage = React.lazy(() => import('../../pages/LandingPage'));
const LogoutPage = React.lazy(() => import('../../pages/LogoutPage'));
const NotFoundPage = React.lazy(() => import('../../pages/NotFoundPage'));
const SignInPage = React.lazy(() => import('../../pages/SignInPage'));
const SignUpPage = React.lazy(() => import('../../pages/SignUpPage'));

export const HOME_ROUTE = '/';
export const REGISTER_ROUTE = '/register';
export const LOGIN_ROUTE = '/login';
export const LOGOUT_ROUTE = '/logout';
export const NOT_FOUND_ROUTE = '/not-found';

const Router = () => {
  const { isAuth } = useAuth();

  return (
    <Suspense fallback={<Loader fullPage />}>
      <BrowserRouter>
        <Switch>
          {/** authenticated routes */}
          <ProtectedRoute
            path={LOGOUT_ROUTE}
            component={LogoutPage}
            condition={isAuth}
            redirect={HOME_ROUTE}
          />

          {/** unauthenticated routes */}
          <ProtectedRoute
            path={LOGIN_ROUTE}
            component={SignInPage}
            condition={!isAuth}
            redirect={HOME_ROUTE}
          />
          <ProtectedRoute
            path={REGISTER_ROUTE}
            component={SignUpPage}
            condition={!isAuth}
            redirect={HOME_ROUTE}
          />

          {/** default routes */}
          <Route
            exact
            path={HOME_ROUTE}
            component={isAuth ? HomePage : LandingPage}
          />
          <Route component={NotFoundPage} />
        </Switch>
      </BrowserRouter>
    </Suspense>
  );
};

export default Router;
