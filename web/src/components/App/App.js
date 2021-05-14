import { CssBaseline } from '@material-ui/core';
import { ErrorBoundary } from '@sentry/react';

import Router from '../Router';
import ThemeProvider from '../Theme';

const App = () => {
  return (
    <ThemeProvider>
      <CssBaseline />
      <ErrorBoundary fallback={'An error has occured'}>
        <Router />
      </ErrorBoundary>
    </ThemeProvider>
  );
};

export default App;
