import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Navigation } from './components/Navigation';
import { TemplateManager } from './components/TemplateManager';
import App from './App';
import './App.css';

export const AppRouter: React.FC = () => {
  const [theme, setTheme] = useState('dark');

  useEffect(() => {
    document.body.setAttribute('data-theme', theme);
  }, [theme]);

  const toggleTheme = () => {
    setTheme((prev) => (prev === 'light' ? 'dark' : 'light'));
  };

  return (
    <Router>
      <div style={{ 
        minHeight: '100vh', 
        display: 'flex', 
        flexDirection: 'column', 
        background: 'var(--background)', 
        color: 'var(--text)', 
        transition: 'background 0.3s, color 0.3s' 
      }}>
        <Navigation theme={theme} toggleTheme={toggleTheme} />
        
        <main style={{ flex: 1 }}>
          <Routes>
            <Route path="/" element={<AppWithoutHeader theme={theme} />} />
            <Route path="/templates" element={<TemplateManager />} />
          </Routes>
        </main>

        <Footer />
      </div>
    </Router>
  );
};

// Wrapper component to pass theme to the original App without header
const AppWithoutHeader: React.FC<{ theme: string }> = ({ theme }) => {
  return <App hideHeader theme={theme} />;
};

const Footer: React.FC = () => (
  <footer style={{ 
    background: 'var(--footer-bg)', 
    color: 'var(--footer-text)', 
    boxShadow: '0 -2px 8px rgba(0,0,0,0.04)', 
    padding: '1rem 2rem', 
    display: 'flex', 
    justifyContent: 'space-between', 
    alignItems: 'center', 
    marginTop: 'auto' 
  }}>
    <span style={{ display: 'flex', alignItems: 'center' }}>
      <svg style={{ height: 24, width: 24, marginRight: 6 }} fill="currentColor" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
        <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.012 8.012 0 0 0 16 8c0-4.42-3.58-8-8-8z" />
      </svg>
    </span>
    <div style={{ display: 'flex', alignItems: 'center', gap: 16 }}>
      <span style={{ color: 'var(--footer-text)', fontWeight: 500, fontSize: 15 }}>
        &copy; {new Date().getFullYear()} Go Initializer. All rights reserved.
      </span>
    </div>
  </footer>
);