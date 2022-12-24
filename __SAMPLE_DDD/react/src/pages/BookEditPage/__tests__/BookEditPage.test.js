import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import AuthProvider from '../../../providers/AuthProvider';
import BookEditPage from '../BookEditPage';

test('render <BookEditPage>', () => {
  render(
    <AuthProvider>
      <MemoryRouter>
        <BookEditPage />
      </MemoryRouter>
    </AuthProvider>,
  );
});
