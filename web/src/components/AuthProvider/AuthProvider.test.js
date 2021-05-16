import React from 'react';
import { render } from '@testing-library/react';

import AuthProvider from './AuthProvider';

const MockComponent = () => {
  return (
    <div>
      <h1>Hello World</h1>
    </div>
  );
};

const customRender = (component) => {
  return render(<AuthProvider>{component}</AuthProvider>);
};

test('renders component with AuthProvider', () => {
  customRender(<MockComponent />);
});
