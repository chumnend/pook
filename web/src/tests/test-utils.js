import { render } from '@testing-library/react';
import { createMemoryHistory } from 'history';
import { Router } from 'react-router-dom';

/**
 * Custom render function to setup components with providers.
 */
export const customRender = (component) => {
  const history = createMemoryHistory();

  render(<Router history={history}>{component}</Router>);

  return { history };
};
