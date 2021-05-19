import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import SignUpPage from './SignUpPage';

test('renders <SignUpPage />', () => {
  render(
    <MemoryRouter>
      <SignUpPage />
    </MemoryRouter>,
  );
});
