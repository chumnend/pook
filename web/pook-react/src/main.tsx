import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';

import App from './components/app/App';
import AuthProvider from './components/providers/AuthProvider';
import './styles.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <App />
    </AuthProvider>
  </StrictMode>,
)
