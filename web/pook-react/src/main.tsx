import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import App from './components/App'
import AuthProvider from './components/providers/AuthProvider'
import './helpers/styles/global.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <App />
    </AuthProvider>
  </StrictMode>,
)
