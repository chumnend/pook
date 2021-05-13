import { render as rtlRender } from '@testing-library/react';
import { createMemoryHistory } from 'history';
import { Router } from 'react-router-dom';

/**
 * Custom render function to setup components with providers.
 */
export const render = (component) => {
  const history = createMemoryHistory();

  rtlRender(<Router history={history}>{component}</Router>);

  return { history };
};
