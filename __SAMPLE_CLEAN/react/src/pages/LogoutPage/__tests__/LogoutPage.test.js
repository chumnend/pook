import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import { AppProviders } from '../../../components/App';
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
