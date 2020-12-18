import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import { ThemeProvider } from '@material-ui/core';
import './index.css';
import App from './App';
import { AuthProvider } from './context/auth';
import theme from './theme';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <AuthProvider>
      <Router>
        <ThemeProvider theme={theme}>
          <App />
        </ThemeProvider>
      </Router>
    </AuthProvider>
  </React.StrictMode>,
  document.getElementById('root'),
);

reportWebVitals();
