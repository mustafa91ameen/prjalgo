import apiClient from './client'

const BASE_PATH = '/api/v1/expenses'

/**
 * Build query string from params object
 * @param {Object} params - Query parameters
 * @returns {string} - Query string
 */
function buildQuery(params = {}) {
  const query = new URLSearchParams()
  if (params.page) query.append('page', params.page)
  if (params.limit) query.append('limit', params.limit)
  if (params.search) query.append('search', params.search)
  if (params.projectId) query.append('projectId', params.projectId)
  if (params.startDate) query.append('startDate', params.startDate)
  if (params.endDate) query.append('endDate', params.endDate)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

export const expensesApi = {
  /**
   * Get all expenses with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {number} [params.projectId] - Filter by project
   * @param {string} [params.startDate] - Filter start date
   * @param {string} [params.endDate] - Filter end date
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Expenses list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${BASE_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single expense by ID
   * @param {number} id - Expense ID
   * @returns {Promise<Object>} - Expense details
   */
  getById(id) {
    return apiClient.get(`${BASE_PATH}/${id}`)
  },

  /**
   * Create a new expense
   * @param {Object} data - Expense data
   * @param {string} data.description - Expense description (required)
   * @param {number} data.amount - Expense amount (required)
   * @param {number} [data.projectId] - Associated project ID
   * @param {string} [data.date] - Expense date
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created expense
   */
  create(data) {
    return apiClient.post(BASE_PATH, data)
  },

  /**
   * Update an existing expense
   * @param {number} id - Expense ID
   * @param {Object} data - Updated expense data
   * @returns {Promise<Object>} - Updated expense
   */
  update(id, data) {
    return apiClient.put(`${BASE_PATH}/${id}`, data)
  },

  /**
   * Delete an expense
   * @param {number} id - Expense ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${BASE_PATH}/${id}`)
  },
}

export default expensesApi
