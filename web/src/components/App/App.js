import { ErrorBoundary } from '@sentry/react';

import AuthProvider from '../AuthProvider';
import Router from '../Router';
import ThemeProvider from '../ThemeProvider';

const App = () => {
  return (
    <AuthProvider>
      <ThemeProvider>
        <ErrorBoundary fallback={'An error has occured'}>
          <Router />
        </ErrorBoundary>
      </ThemeProvider>
    </AuthProvider>
  );
};

export default App;
