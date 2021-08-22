import { ErrorBoundary } from '@sentry/react';
import PropTypes from 'prop-types';

import AuthProvider from '../AuthProvider';
import Router from '../Router';
import ThemeProvider from '../ThemeProvider';

export const AppProviders = ({ children }) => {
  return (
    <AuthProvider>
      <ThemeProvider>{children}</ThemeProvider>
    </AuthProvider>
  );
};

AppProviders.propTypes = {
  children: PropTypes.node.isRequired,
};

const App = () => {
  return (
    <AppProviders>
      <ErrorBoundary fallback={'An error has occured'}>
        <Router />
      </ErrorBoundary>
    </AppProviders>
  );
};

export default App;
