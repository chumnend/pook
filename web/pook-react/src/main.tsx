import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import App from './components/App.tsx'
import './helpers/global.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
