import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import AuthProvider from '../../../providers/AuthProvider';
import LandingPage from '../LandingPage';

test('render <LandingPage>', () => {
  render(
    <AuthProvider>
      <MemoryRouter>
        <LandingPage />
      </MemoryRouter>
    </AuthProvider>,
  );
});
