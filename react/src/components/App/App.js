import { ErrorBoundary } from '@sentry/react';
import PropTypes from 'prop-types';

import AuthProvider from '../../providers/AuthProvider';
import ThemeProvider from '../../providers/ThemeProvider';
import Router from '../Router';

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
