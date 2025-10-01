import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { AppRouter } from './AppRouter';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <AppRouter />
  </React.StrictMode>
);

// Performance monitoring can be enabled by passing a function to reportWebVitals
reportWebVitals();
