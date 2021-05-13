import { CssBaseline } from '@material-ui/core';
import { MuiThemeProvider } from '@material-ui/core/styles';
import { ErrorBoundary } from '@sentry/react';

import theme from '../../common/theme';

const App = () => {
  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <ErrorBoundary fallback={'An error has occured'}>
        <h1>Hello World</h1>
      </ErrorBoundary>
    </MuiThemeProvider>
  );
};

export default App;
