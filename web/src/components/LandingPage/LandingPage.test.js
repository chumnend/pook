import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import LandingPage from './LandingPage';

it('render <LandingPage>', () => {
  render(
    <MemoryRouter>
      <LandingPage />
    </MemoryRouter>,
  );
});
