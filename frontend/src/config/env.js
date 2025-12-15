export const config = {
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:5000',
  apiTimeout: parseInt(import.meta.env.VITE_API_TIMEOUT) || 30000,
  appTitle: import.meta.env.VITE_APP_TITLE || 'Project Management System',
  appEnv: import.meta.env.VITE_APP_ENV || 'development',
}
