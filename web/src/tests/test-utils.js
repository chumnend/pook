import { render } from '@testing-library/react';
import { createMemoryHistory } from 'history';
import { Router } from 'react-router-dom';

import { AuthProvider } from '../services/context/auth';

/**
 * Custom render function to setup components with providers.
 */
export const customRender = (component) => {
  const history = createMemoryHistory();

  render(
    <AuthProvider>
      <Router history={history}>{component}</Router>
    </AuthProvider>,
  );

  return { history };
};
