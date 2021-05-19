import { render } from '@testing-library/react';

import AuthProvider from '../AuthProvider';
import Router from './Router';

test('renders <Router />', () => {
  render(
    <AuthProvider>
      <Router />
    </AuthProvider>,
  );
});
