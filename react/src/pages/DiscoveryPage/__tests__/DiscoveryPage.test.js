import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import AuthProvider from '../../../providers/AuthProvider';
import DiscoveryPage from '../DiscoveryPage';

test('render <DiscoveryPage>', () => {
  render(
    <AuthProvider>
      <MemoryRouter>
        <DiscoveryPage />
      </MemoryRouter>
    </AuthProvider>,
  );
});
