import { apiFetch } from './client'

/**
 * Get dashboard stats (total projects, active, completed, engineers)
 */
export async function getDashboardStats() {
  return apiFetch('/dashboard/stats')
}

/**
 * Get financial stats (total income, total expenses)
 * @param {Object} params - Query parameters
 * @param {string} params.period - 'all' or 'month'
 * @param {string} params.month - Month in YYYY-MM format (required when period='month')
 */
export async function getFinancialStats(params = {}) {
  const query = new URLSearchParams()
  if (params.period) query.append('period', params.period)
  if (params.month) query.append('month', params.month)
  const queryString = query.toString()
  return apiFetch(`/dashboard/financial${queryString ? `?${queryString}` : ''}`)
}

/**
 * Get project progress
 * @param {Object} params - Query parameters
 * @param {string} params.status - 'pending', 'in_progress', 'completed', or 'all'
 * @param {number} params.limit - Max projects to return (1-20)
 */
export async function getProjectProgress(params = {}) {
  const query = new URLSearchParams()
  if (params.status) query.append('status', params.status)
  if (params.limit) query.append('limit', params.limit.toString())
  const queryString = query.toString()
  return apiFetch(`/dashboard/project-progress${queryString ? `?${queryString}` : ''}`)
}

/**
 * Get task summary for a specific month
 * @param {string} month - Month in YYYY-MM format
 */
export async function getTaskSummary(month) {
  return apiFetch(`/dashboard/task-summary?month=${month}`)
}

/**
 * Get recent activities (audit logs)
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number
 * @param {number} params.limit - Items per page (max 100)
 */
export async function getActivities(params = {}) {
  const query = new URLSearchParams()
  if (params.page) query.append('page', params.page.toString())
  if (params.limit) query.append('limit', params.limit.toString())
  const queryString = query.toString()
  return apiFetch(`/dashboard/activities${queryString ? `?${queryString}` : ''}`)
}
