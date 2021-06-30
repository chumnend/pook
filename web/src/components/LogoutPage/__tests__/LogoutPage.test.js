import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import { AppProviders } from '../../App';
import LogoutPage from '../LogoutPage';

test('renders <LogoutPage />', () => {
  render(
    <MemoryRouter>
      <AppProviders>
        <LogoutPage />
      </AppProviders>
    </MemoryRouter>,
  );
});
