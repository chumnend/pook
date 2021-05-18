import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import AuthProvider from '../AuthProvider';
import NotFoundPage from './NotFoundPage';

it('render <NotFoundPage>', () => {
  render(
    <AuthProvider>
      <MemoryRouter>
        <NotFoundPage />
      </MemoryRouter>
    </AuthProvider>,
  );
});
