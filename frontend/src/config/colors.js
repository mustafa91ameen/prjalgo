/**
 * Centralized Color Configuration
 *
 * This file is the SINGLE SOURCE OF TRUTH for all colors.
 * Both Vuetify (vuetify.js) and CSS (colors.css) should use these values.
 *
 * To change a color:
 * 1. Update it here
 * 2. Run the app - Vuetify components will update automatically
 * 3. Update colors.css with the same value (or use CSS custom properties)
 */

export const colors = {
  // Primary - Electric Indigo (Crystal Indigo Theme)
  primary: '#4f46e5',           // Indigo 600 - Main actions, active states
  'primary-light': '#6366f1',   // Indigo 500
  'primary-dark': '#3730a3',    // Indigo 800

  // Secondary - Soft Sky
  secondary: '#64748b',         // Slate 500
  'secondary-light': '#94a3b8', // Slate 400
  'secondary-dark': '#475569',  // Slate 600

  // Success - Green (Pastel style per Crystal Indigo)
  success: '#166534',           // Green 800 (Text)
  'success-light': '#dcfce7',   // Green 100 (Background)
  'success-dark': '#14532d',    // Green 900

  // Warning - Amber (Pastel style)
  warning: '#92400e',           // Amber 800 (Text)
  'warning-light': '#fef3c7',   // Amber 100 (Background)
  'warning-dark': '#78350f',    // Amber 900

  // Error - Red (Pastel style)
  error: '#991b1b',             // Red 800 (Text)
  'error-light': '#fee2e2',     // Red 100 (Background)
  'error-dark': '#7f1d1d',      // Red 900

  // Info - Blue (Pastel style)
  info: '#1e40af',              // Blue 800 (Text)
  'info-light': '#dbeafe',      // Blue 100 (Background)
  'info-dark': '#1e3a8a',       // Blue 900

  // Background & Surface (Crystal Indigo)
  background: '#f8fafc',        // Slate 50 - App background
  surface: '#ffffff',           // Pure Crystal - Cards, Sidebar, Modals
}

export default colors
