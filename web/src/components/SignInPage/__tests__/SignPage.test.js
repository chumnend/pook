import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import SignInPage from '../SignInPage';

test('renders <SignInPage />', () => {
  render(
    <MemoryRouter>
      <SignInPage />
    </MemoryRouter>,
  );
});
