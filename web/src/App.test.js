import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import App from './App';
import { AuthProvider } from './context/auth';

test('render <App />', () => {
  render(
    <AuthProvider>
      <MemoryRouter initialEntries={['/']}>
        <App />
      </MemoryRouter>
    </AuthProvider>,
  )

  const el = screen.getByText(/Landing/i);
  expect(el).toBeInTheDocument();
})
