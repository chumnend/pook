import { render, screen } from '@testing-library/react';
import App from './App';
import { BrowserRouter } from 'react-router-dom';

test('renders <App />', () => {
  render(
    <BrowserRouter>
      <App />
    </BrowserRouter>
  );

  const el = screen.getByText(/Pook/i);
  expect(el).toBeInTheDocument();
});
