import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import App from './App';

test('renders home page', () => {
  render(
    <MemoryRouter initialEntries={['/home']}>
      <App />
    </MemoryRouter>,
  );

  const el = screen.getByText(/Home/i);
  expect(el).toBeInTheDocument();
});

test('renders landing page', () => {
  render(
    <MemoryRouter initialEntries={['/']}>
      <App />
    </MemoryRouter>,
  );

  const el = screen.getByText(/Landing/i);
  expect(el).toBeInTheDocument();
});

test('renders not found page', () => {
  render(
    <MemoryRouter initialEntries={['/not-found']}>
      <App />
    </MemoryRouter>,
  );

  const el = screen.getByText(/Page not found/i);
  expect(el).toBeInTheDocument();
});
