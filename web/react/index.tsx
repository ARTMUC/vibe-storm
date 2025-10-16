import { createRoot } from 'react-dom/client';
import { AuthStatus } from './AuthStatus';
import { APITesting } from './APITesting';

// API Status Checker (vanilla JS, no React needed)
const checkAPIStatus = async () => {
  const indicator = document.getElementById('api-status-indicator');
  const message = document.getElementById('api-status-message');
  
  if (!indicator || !message) return;

  try {
    const response = await fetch('/api/v1/health');
    const data = await response.json();
    
    if (data.status === 'healthy') {
      indicator.className = 'status-indicator success';
      message.textContent = `API is healthy (${data.service} v${data.version})`;
    } else {
      indicator.className = 'status-indicator error';
      message.textContent = 'API returned unhealthy status';
    }
  } catch (error) {
    indicator.className = 'status-indicator error';
    message.textContent = 'Failed to connect to API';
  }
};

// Initialize when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
  // Mount AuthStatus component
  const authStatusRoot = document.getElementById('auth-status-root');
  if (authStatusRoot) {
    const root = createRoot(authStatusRoot);
    root.render(<AuthStatus />);
  }

  // Mount APITesting component
  const apiTestingRoot = document.getElementById('api-testing-root');
  if (apiTestingRoot) {
    const root = createRoot(apiTestingRoot);
    root.render(<APITesting />);
  }

  // Check API status
  checkAPIStatus();
});

// Export for potential external use
export { AuthStatus, APITesting };
