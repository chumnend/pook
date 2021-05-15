import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import ProtectedRoute from './ProtectedRoute';

test('renders <ProtectedRoute /> with true condition', () => {
  render(
    <MemoryRouter>
      <ProtectedRoute path="/" condition={true} redirect={'/redirect'} />
    </MemoryRouter>,
  );
});

test('renders <ProtectedRoute /> with false condition', () => {
  render(
    <MemoryRouter>
      <ProtectedRoute path="/" condition={false} redirect={'/redirect'} />
    </MemoryRouter>,
  );
});
