/**
 * API Module Barrel Export
 * Centralizes all API modules for easy imports
 */

// Core client
export { default as apiClient, TOKEN_KEY, REFRESH_TOKEN_KEY } from './client'

// Auth
export { authApi } from './auth'

// Projects
export { projectsApi } from './projects'

// Debtors
export { debtorsApi } from './debtors'

// Expenses
export { expensesApi } from './expenses'

// Income
export { incomeApi } from './income'

// Users & User Roles
export { usersApi, userRolesApi } from './users'

// Team Members
export { teamsApi } from './teams'

// Workdays, Materials, Labor, Equipment
export {
  workdaysApi,
  workdayMaterialsApi,
  workdayLaborApi,
  workdayEquipmentApi,
} from './workdays'

// Roles, Pages, Role-Pages
export { rolesApi, pagesApi, rolePagesApi } from './roles'

// Work Categories & Subcategories
export { workCategoriesApi, workSubcategoriesApi } from './categories'
