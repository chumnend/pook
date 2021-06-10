import { render } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';

import Header from '../Header';

it('render <Header>', () => {
  render(
    <MemoryRouter>
        <Header />
    </MemoryRouter>,
  );
});
