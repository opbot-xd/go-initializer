import React from 'react';
import { Link, useLocation } from 'react-router-dom';

interface NavigationProps {
  theme: string;
  toggleTheme: () => void;
}

export const Navigation: React.FC<NavigationProps> = ({ theme, toggleTheme }) => {
  const location = useLocation();

  const navItems = [
    { path: '/', label: 'Generator', icon: '🚀' },
    { path: '/templates', label: 'Templates', icon: '📝' },
  ];

  return (
    <header style={{ 
      background: 'var(--navbar-bg)', 
      boxShadow: '0 2px 8px rgba(0,0,0,0.04)', 
      padding: '1rem 2rem', 
      display: 'flex', 
      justifyContent: 'space-between', 
      alignItems: 'center', 
      color: 'var(--navbar-text)' 
    }}>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <Link to="/" style={{ textDecoration: 'none', display: 'flex', alignItems: 'center' }}>
          <div style={{ 
            width: 40, 
            height: 40, 
            background: 'linear-gradient(135deg, #ffd700, #ffed4e)', 
            borderRadius: '50%', 
            display: 'flex', 
            alignItems: 'center', 
            justifyContent: 'center', 
            marginRight: 16,
            fontSize: '20px'
          }}>
            🐹
          </div>
          <h1 style={{ 
            fontSize: 28, 
            fontWeight: 800, 
            color: 'var(--navbar-text)', 
            letterSpacing: 0.5,
            margin: 0
          }}>
            go <span style={{ color: '#ffd700' }}>initializer</span>
          </h1>
        </Link>
      </div>

      <nav style={{ display: 'flex', alignItems: 'center', gap: '2rem' }}>
        {navItems.map((item) => (
          <Link
            key={item.path}
            to={item.path}
            style={{
              textDecoration: 'none',
              color: location.pathname === item.path ? '#ffd700' : 'var(--navbar-text)',
              fontWeight: location.pathname === item.path ? 700 : 500,
              fontSize: 16,
              display: 'flex',
              alignItems: 'center',
              gap: '0.5rem',
              padding: '0.5rem 1rem',
              borderRadius: '8px',
              background: location.pathname === item.path ? 'rgba(255, 215, 0, 0.1)' : 'transparent',
              transition: 'all 0.2s ease',
            }}
          >
            <span>{item.icon}</span>
            {item.label}
          </Link>
        ))}
      </nav>

      <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
        <button 
          style={{ 
            padding: 8, 
            borderRadius: '50%', 
            background: 'none', 
            border: 'none', 
            color: '#666', 
            cursor: 'pointer', 
            fontSize: 22 
          }} 
          title="Toggle theme" 
          onClick={toggleTheme}
        >
          {theme === 'light' ? '🌙' : '☀️'}
        </button>
      </div>
    </header>
  );
};