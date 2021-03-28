import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import App from './App';

test('renders home page', () => {
  render(
    <MemoryRouter initialEntries={['/']}>
      <App />
    </MemoryRouter>,
  );

  const el = screen.getByText(/Pook/i);
  expect(el).toBeInTheDocument();
});

test('renders home page', () => {
  render(
    <MemoryRouter initialEntries={['/not-found']}>
      <App />
    </MemoryRouter>,
  );

  const el = screen.getByText(/Page not found/i);
  expect(el).toBeInTheDocument();
});

