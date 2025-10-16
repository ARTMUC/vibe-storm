import React, { useState } from 'react';

type TabType = 'signup' | 'signin' | 'refresh' | 'profile';

interface FormData {
  email: string;
  username: string;
  password: string;
  firstName: string;
  lastName: string;
  refreshToken: string;
}

export const APITesting: React.FC = () => {
  const [activeTab, setActiveTab] = useState<TabType>('signup');
  const [response, setResponse] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState<FormData>({
    email: '',
    username: '',
    password: '',
    firstName: '',
    lastName: '',
    refreshToken: ''
  });

  const handleInputChange = (field: keyof FormData, value: string) => {
    setFormData(prev => ({ ...prev, [field]: value }));
  };

  const makeAPICall = async (endpoint: string, data: any) => {
    setLoading(true);
    try {
      const response = await fetch(`/api/v1${endpoint}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      const result = await response.json();

      if (response.ok) {
        setResponse(result);
        // If it's a successful auth response, store tokens
        if (result.access_token && result.refresh_token) {
          localStorage.setItem('accessToken', result.access_token);
          localStorage.setItem('refreshToken', result.refresh_token);
          localStorage.setItem('tokenExpiry', result.expires_at);
          // Notify other components about auth update
          window.dispatchEvent(new Event('auth-updated'));
        }
      } else {
        setResponse({ error: result.message || 'API call failed' });
      }
    } catch (error: any) {
      setResponse({ error: error.message || 'Network error' });
    } finally {
      setLoading(false);
    }
  };

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    await makeAPICall('/auth/signup', {
      email: formData.email,
      username: formData.username,
      password: formData.password,
      first_name: formData.firstName,
      last_name: formData.lastName
    });
  };

  const handleSignin = async (e: React.FormEvent) => {
    e.preventDefault();
    await makeAPICall('/auth/signin', {
      email: formData.email,
      password: formData.password
    });
  };

  const handleRefresh = async (e: React.FormEvent) => {
    e.preventDefault();
    await makeAPICall('/auth/refresh', {
      refresh_token: formData.refreshToken
    });
  };

  const getProfile = async () => {
    const token = localStorage.getItem('accessToken');
    if (!token) {
      setResponse({ error: 'No access token found. Please sign in first.' });
      return;
    }

    setLoading(true);
    try {
      const response = await fetch('/api/v1/auth/me', {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      const result = await response.json();
      if (response.ok) {
        setResponse(result);
      } else {
        setResponse({ error: result.message || 'Failed to get profile' });
      }
    } catch (error: any) {
      setResponse({ error: error.message || 'Network error' });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="api-testing-container">
      <div className="api-tabs">
        <button
          className={`tab-btn ${activeTab === 'signup' ? 'active' : ''}`}
          onClick={() => setActiveTab('signup')}
        >
          Signup
        </button>
        <button
          className={`tab-btn ${activeTab === 'signin' ? 'active' : ''}`}
          onClick={() => setActiveTab('signin')}
        >
          Signin
        </button>
        <button
          className={`tab-btn ${activeTab === 'refresh' ? 'active' : ''}`}
          onClick={() => setActiveTab('refresh')}
        >
          Refresh Token
        </button>
        <button
          className={`tab-btn ${activeTab === 'profile' ? 'active' : ''}`}
          onClick={() => setActiveTab('profile')}
        >
          Get Profile
        </button>
      </div>

      <div className="tab-content">
        {activeTab === 'signup' && (
          <form onSubmit={handleSignup}>
            <div className="form-group">
              <label>Email:</label>
              <input
                type="email"
                value={formData.email}
                onChange={(e) => handleInputChange('email', e.target.value)}
                required
                placeholder="user@example.com"
              />
            </div>
            <div className="form-group">
              <label>Username:</label>
              <input
                type="text"
                value={formData.username}
                onChange={(e) => handleInputChange('username', e.target.value)}
                required
                placeholder="johndoe"
              />
            </div>
            <div className="form-group">
              <label>Password:</label>
              <input
                type="password"
                value={formData.password}
                onChange={(e) => handleInputChange('password', e.target.value)}
                required
                placeholder="SecurePass123!"
              />
            </div>
            <div className="form-group">
              <label>First Name:</label>
              <input
                type="text"
                value={formData.firstName}
                onChange={(e) => handleInputChange('firstName', e.target.value)}
                required
                placeholder="John"
              />
            </div>
            <div className="form-group">
              <label>Last Name:</label>
              <input
                type="text"
                value={formData.lastName}
                onChange={(e) => handleInputChange('lastName', e.target.value)}
                required
                placeholder="Doe"
              />
            </div>
            <button type="submit" className="btn btn-primary" disabled={loading}>
              {loading ? 'Signing up...' : 'Sign Up'}
            </button>
          </form>
        )}

        {activeTab === 'signin' && (
          <form onSubmit={handleSignin}>
            <div className="form-group">
              <label>Email:</label>
              <input
                type="email"
                value={formData.email}
                onChange={(e) => handleInputChange('email', e.target.value)}
                required
                placeholder="user@example.com"
              />
            </div>
            <div className="form-group">
              <label>Password:</label>
              <input
                type="password"
                value={formData.password}
                onChange={(e) => handleInputChange('password', e.target.value)}
                required
                placeholder="password"
              />
            </div>
            <button type="submit" className="btn btn-primary" disabled={loading}>
              {loading ? 'Signing in...' : 'Sign In'}
            </button>
          </form>
        )}

        {activeTab === 'refresh' && (
          <form onSubmit={handleRefresh}>
            <div className="form-group">
              <label>Refresh Token:</label>
              <textarea
                value={formData.refreshToken}
                onChange={(e) => handleInputChange('refreshToken', e.target.value)}
                required
                placeholder="Paste your refresh token here"
              />
            </div>
            <button type="submit" className="btn btn-primary" disabled={loading}>
              {loading ? 'Refreshing...' : 'Refresh Token'}
            </button>
          </form>
        )}

        {activeTab === 'profile' && (
          <div className="profile-section">
            <p>Click the button to fetch your profile information using your stored access token.</p>
            <button
              type="button"
              className="btn btn-primary"
              onClick={getProfile}
              disabled={loading}
            >
              {loading ? 'Loading...' : 'Get Profile'}
            </button>
          </div>
        )}
      </div>

      {response && (
        <div className="response-display">
          <h3>API Response:</h3>
          <pre>{JSON.stringify(response, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};
