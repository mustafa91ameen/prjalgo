/**
 * Application constants and enums
 * Centralizes all constant values used across the application
 */

// ==================== PROJECT STATUS ====================

export const PROJECT_STATUS = {
  PENDING: 'pending',
  IN_PROGRESS: 'in_progress',
  COMPLETED: 'completed',
  ON_HOLD: 'on_hold',
  CANCELLED: 'cancelled',
}

export const PROJECT_STATUS_LABELS = {
  [PROJECT_STATUS.PENDING]: 'قيد الانتظار',
  [PROJECT_STATUS.IN_PROGRESS]: 'قيد التنفيذ',
  [PROJECT_STATUS.COMPLETED]: 'مكتمل',
  [PROJECT_STATUS.ON_HOLD]: 'معلق',
  [PROJECT_STATUS.CANCELLED]: 'ملغي',
}

export const PROJECT_STATUS_COLORS = {
  [PROJECT_STATUS.PENDING]: 'warning',
  [PROJECT_STATUS.IN_PROGRESS]: 'primary',
  [PROJECT_STATUS.COMPLETED]: 'success',
  [PROJECT_STATUS.ON_HOLD]: 'secondary',
  [PROJECT_STATUS.CANCELLED]: 'danger',
}

// ==================== USER STATUS ====================

export const USER_STATUS = {
  ACTIVE: 'active',
  INACTIVE: 'inactive',
  SUSPENDED: 'suspended',
}

export const USER_STATUS_LABELS = {
  [USER_STATUS.ACTIVE]: 'نشط',
  [USER_STATUS.INACTIVE]: 'غير نشط',
  [USER_STATUS.SUSPENDED]: 'معلق',
}

export const USER_STATUS_COLORS = {
  [USER_STATUS.ACTIVE]: 'success',
  [USER_STATUS.INACTIVE]: 'secondary',
  [USER_STATUS.SUSPENDED]: 'danger',
}

// ==================== WORKDAY STATUS ====================

export const WORKDAY_STATUS = {
  PENDING: 'pending',
  IN_PROGRESS: 'in_progress',
  COMPLETED: 'completed',
}

export const WORKDAY_STATUS_LABELS = {
  [WORKDAY_STATUS.PENDING]: 'قيد الانتظار',
  [WORKDAY_STATUS.IN_PROGRESS]: 'قيد التنفيذ',
  [WORKDAY_STATUS.COMPLETED]: 'مكتمل',
}

export const WORKDAY_STATUS_COLORS = {
  [WORKDAY_STATUS.PENDING]: 'warning',
  [WORKDAY_STATUS.IN_PROGRESS]: 'primary',
  [WORKDAY_STATUS.COMPLETED]: 'success',
}

// ==================== PERMISSIONS ====================

export const PERMISSIONS = {
  READ: 'read',
  WRITE: 'write',
  DELETE: 'delete',
}

export const PERMISSION_LABELS = {
  [PERMISSIONS.READ]: 'قراءة',
  [PERMISSIONS.WRITE]: 'كتابة',
  [PERMISSIONS.DELETE]: 'حذف',
}

// ==================== PAGINATION ====================

export const PAGINATION = {
  DEFAULT_PAGE: 1,
  DEFAULT_LIMIT: 10,
  PAGE_SIZE_OPTIONS: [10, 25, 50, 100],
}

// ==================== DATE FORMATS ====================

export const DATE_FORMATS = {
  DISPLAY: 'DD/MM/YYYY',
  API: 'YYYY-MM-DD',
  DATETIME_DISPLAY: 'DD/MM/YYYY HH:mm',
  DATETIME_API: 'YYYY-MM-DDTHH:mm:ss',
}

// ==================== API ENDPOINTS (for reference) ====================

export const API_ENDPOINTS = {
  // Auth
  AUTH_LOGIN: '/api/v1/auth/login',
  AUTH_LOGOUT: '/api/v1/auth/logout',
  AUTH_REFRESH: '/api/v1/auth/refresh',
  AUTH_PAGES: '/api/v1/auth/pages',

  // Projects
  PROJECTS: '/api/v1/projects',

  // Debtors
  DEBTORS: '/api/v1/debtors',

  // Expenses
  EXPENSES: '/api/v1/expenses',

  // Income
  INCOME: '/api/v1/income',

  // Users
  USERS: '/api/v1/users',
  USER_ROLES: '/api/v1/user-roles',

  // Team Members
  TEAM_MEMBERS: '/api/v1/team-members',

  // Workdays
  WORKDAYS: '/api/v1/workdays',
  WORKDAY_MATERIALS: '/api/v1/workday-materials',
  WORKDAY_LABOR: '/api/v1/workday-labor',
  WORKDAY_EQUIPMENT: '/api/v1/workday-equipment',

  // Roles & Pages
  ROLES: '/api/v1/roles',
  PAGES: '/api/v1/pages',
  ROLE_PAGES: '/api/v1/role-pages',

  // Categories
  WORK_CATEGORIES: '/api/v1/work-categories',
  WORK_SUBCATEGORIES: '/api/v1/work-subcategories',
}

// ==================== ROUTES (Frontend) ====================

export const ROUTES = {
  HOME: '/',
  LOGIN: '/login',
  DASHBOARD: '/dashboard',
  PROJECTS: '/projects',
  PROJECT_DETAILS: '/projects/:id',
  DEBTORS: '/debtors',
  EXPENSES: '/expenses',
  INCOME: '/income',
  USERS: '/users',
  ROLES: '/roles',
  TEAM_MEMBERS: '/team-members',
  WORKDAYS: '/workdays',
  CATEGORIES: '/categories',
  SETTINGS: '/settings',
}

// ==================== MODALS ====================

export const MODALS = {
  // Project modals
  PROJECT_CREATE: 'project-create',
  PROJECT_EDIT: 'project-edit',
  PROJECT_DELETE: 'project-delete',

  // Debtor modals
  DEBTOR_CREATE: 'debtor-create',
  DEBTOR_EDIT: 'debtor-edit',
  DEBTOR_DELETE: 'debtor-delete',

  // Expense modals
  EXPENSE_CREATE: 'expense-create',
  EXPENSE_EDIT: 'expense-edit',
  EXPENSE_DELETE: 'expense-delete',

  // User modals
  USER_CREATE: 'user-create',
  USER_EDIT: 'user-edit',
  USER_DELETE: 'user-delete',
  USER_PASSWORD: 'user-password',

  // Confirmation modals
  CONFIRM_DELETE: 'confirm-delete',
  CONFIRM_ACTION: 'confirm-action',
}

// ==================== LOCAL STORAGE KEYS ====================

export const STORAGE_KEYS = {
  ACCESS_TOKEN: 'accessToken',
  REFRESH_TOKEN: 'refreshToken',
  SIDEBAR_COLLAPSED: 'sidebarCollapsed',
  DARK_MODE: 'darkMode',
  LANGUAGE: 'language',
  USER_PREFERENCES: 'userPreferences',
}

// ==================== NOTIFICATION TYPES ====================

export const NOTIFICATION_TYPES = {
  SUCCESS: 'success',
  ERROR: 'error',
  WARNING: 'warning',
  INFO: 'info',
}

// ==================== FILE TYPES ====================

export const FILE_TYPES = {
  IMAGE: ['image/jpeg', 'image/png', 'image/gif', 'image/webp'],
  DOCUMENT: ['application/pdf', 'application/msword', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'],
  SPREADSHEET: ['application/vnd.ms-excel', 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'],
}

export const MAX_FILE_SIZE = 10 * 1024 * 1024 // 10MB

// ==================== COLORS ====================

export const COLORS = {
  PRIMARY: '#3b82f6',
  SECONDARY: '#64748b',
  SUCCESS: '#22c55e',
  DANGER: '#ef4444',
  WARNING: '#f59e0b',
  INFO: '#06b6d4',
}

// ==================== BREAKPOINTS ====================

export const BREAKPOINTS = {
  SM: 640,
  MD: 768,
  LG: 1024,
  XL: 1280,
  XXL: 1536,
}

export default {
  PROJECT_STATUS,
  PROJECT_STATUS_LABELS,
  PROJECT_STATUS_COLORS,
  USER_STATUS,
  USER_STATUS_LABELS,
  USER_STATUS_COLORS,
  WORKDAY_STATUS,
  WORKDAY_STATUS_LABELS,
  WORKDAY_STATUS_COLORS,
  PERMISSIONS,
  PERMISSION_LABELS,
  PAGINATION,
  DATE_FORMATS,
  API_ENDPOINTS,
  ROUTES,
  MODALS,
  STORAGE_KEYS,
  NOTIFICATION_TYPES,
  FILE_TYPES,
  MAX_FILE_SIZE,
  COLORS,
  BREAKPOINTS,
}
