import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import { AppProviders } from '../../../components/App';
import HomePage from '../HomePage';

test('renders <HomePage />', () => {
  render(
    <MemoryRouter>
      <AppProviders>
        <HomePage />
      </AppProviders>
    </MemoryRouter>,
  );
});
