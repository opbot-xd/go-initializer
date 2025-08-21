import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';


function App() {
  const [theme, setTheme] = useState('light');

  useEffect(() => {
    document.body.setAttribute('data-theme', theme);
  }, [theme]);

  const toggleTheme = () => {
    setTheme((prev) => (prev === 'light' ? 'dark' : 'light'));
  };

  return (
    <div
      className="App"
      style={{
        minHeight: '100vh',
        display: 'flex',
        flexDirection: 'column',
        background: 'var(--background)',
        color: 'var(--text)',
        transition: 'background 0.3s, color 0.3s',
      }}
    >
      {/* Navbar Header */}
      <nav style={{
        background: 'var(--navbar-bg)',
        boxShadow: '0 2px 8px rgba(0,0,0,0.04)',
        display: 'flex',
        justifyContent: 'center',
      }}>
        <div style={{
          width: '100%',
          maxWidth: 700,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          padding: '0.75rem 2rem',
        }}>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <img src={logo} className="App-logo" alt="logo" style={{ height: 40, marginRight: 16 }} />
            <span style={{ color: 'var(--navbar-text)', fontSize: 24, fontWeight: 600, letterSpacing: 1 }}>Go Initializer</span>
          </div>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <a href="/" style={{ color: 'var(--navbar-link)', textDecoration: 'none', margin: '0 1rem', fontWeight: 500 }}>Home</a>
            <a href="#features" style={{ color: 'var(--navbar-link)', textDecoration: 'none', margin: '0 1rem', fontWeight: 500 }}>Features</a>
            <a href="#about" style={{ color: 'var(--navbar-link)', textDecoration: 'none', margin: '0 1rem', fontWeight: 500 }}>About</a>
            <button
              onClick={toggleTheme}
              style={{
                marginLeft: 24,
                padding: '0.4rem 1rem',
                borderRadius: 6,
                border: 'none',
                background: 'var(--theme-btn-bg)',
                color: 'var(--theme-btn-text)',
                cursor: 'pointer',
                fontWeight: 500,
                fontSize: 15,
                transition: 'background 0.2s, color 0.2s',
              }}
              aria-label="Toggle theme"
            >
              {theme === 'light' ? '🌙 Dark' : '☀️ Light'}
            </button>
          </div>
        </div>
      </nav>

      {/* Content */}
      <main
        style={{
          flex: 1,
          background: 'var(--content-bg)',
          color: 'var(--text)',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'stretch',
          padding: 0,
          minHeight: 0,
        }}
      >
        <div
          style={{
            width: '100%',
            height: '100%',
            maxWidth: 700,
            padding: 0,
            background: 'var(--center-card-bg, #fffbe6)',
            boxShadow: '0 4px 24px rgba(0,0,0,0.08)',
            transition: 'background 0.3s',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
            minHeight: '100%',
          }}
        >
          <h2>Welcome to the Go Initializer SPA</h2>
          <p>
            This is your starting point for building a modern single page application with React and Go.
          </p>
          {/* Add your main content here */}
        </div>
      </main>

      {/* Footer */}
      <footer style={{ background: 'var(--footer-bg)', color: 'var(--footer-text)', textAlign: 'center', padding: '1rem 0' }}>
        <small>&copy; {new Date().getFullYear()} Go Initializer. All rights reserved.</small>
      </footer>
    </div>
  );
}

export default App;
