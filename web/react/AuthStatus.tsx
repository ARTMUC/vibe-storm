import React, { useState, useEffect } from 'react';

interface AuthState {
  isAuthenticated: boolean;
  userEmail: string;
  tokenExpiry: string;
}

export const AuthStatus: React.FC = () => {
  const [authState, setAuthState] = useState<AuthState>({
    isAuthenticated: false,
    userEmail: '',
    tokenExpiry: ''
  });

  useEffect(() => {
    checkAuthStatus();
    // Listen for auth changes from other components
    window.addEventListener('auth-updated', checkAuthStatus);
    return () => window.removeEventListener('auth-updated', checkAuthStatus);
  }, []);

  const checkAuthStatus = () => {
    const accessToken = localStorage.getItem('accessToken');
    const expiryStr = localStorage.getItem('tokenExpiry');

    if (accessToken && expiryStr) {
      const expiryDate = new Date(expiryStr);
      if (expiryDate > new Date()) {
        try {
          const tokenParts = accessToken.split('.');
          if (tokenParts.length >= 2 && tokenParts[1]) {
            const payload = JSON.parse(atob(tokenParts[1]));
            setAuthState({
              isAuthenticated: true,
              userEmail: payload.email || 'Unknown',
              tokenExpiry: expiryDate.toLocaleString()
            });
          } else {
            clearAuth();
          }
        } catch (e) {
          clearAuth();
        }
      } else {
        clearAuth();
      }
    } else {
      clearAuth();
    }
  };

  const clearAuth = () => {
    setAuthState({
      isAuthenticated: false,
      userEmail: '',
      tokenExpiry: ''
    });
  };

  const logout = () => {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    localStorage.removeItem('tokenExpiry');
    clearAuth();
    // Notify other components
    window.dispatchEvent(new Event('auth-updated'));
  };

  return (
    <div className="auth-info">
      <div className="auth-indicator">
        <span className={`status-dot ${authState.isAuthenticated ? 'authenticated' : ''}`}></span>
        <span>{authState.isAuthenticated ? 'Authenticated' : 'Not authenticated'}</span>
      </div>
      {authState.isAuthenticated && (
        <div className="user-info">
          <p><strong>User:</strong> {authState.userEmail}</p>
          <p><strong>Token expires:</strong> {authState.tokenExpiry}</p>
          <button className="btn btn-small" onClick={logout}>Logout</button>
        </div>
      )}
    </div>
  );
};
