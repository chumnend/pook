import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import AuthProvider from '../AuthProvider';
import LandingPage from './LandingPage';

it('render <LandingPage>', () => {
  render(
    <AuthProvider>
      <MemoryRouter>
        <LandingPage />
      </MemoryRouter>
    </AuthProvider>,
  );
});
