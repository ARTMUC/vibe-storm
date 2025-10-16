// Legacy vanilla JS file - keeping minimal utilities for backward compatibility
// All interactive functionality has been moved to React components

// Simple utility functions that might be needed for non-React parts
class LegacyUtils {
	static showNotification(message: string, type: 'success' | 'error' | 'info' = 'info'): void {
		console.log(`[${type.toUpperCase()}] ${message}`);
		// React components handle notifications now
	}

	static log(message: string): void {
		console.log(`[VibeStorm] ${message}`);
	}
}

// Export for any external usage
export { LegacyUtils };
