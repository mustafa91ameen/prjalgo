import apiClient from './client'

const BASE_PATH = '/api/v1/debtors'

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
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

export const debtorsApi = {
  /**
   * Get all debtors with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Debtors list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${BASE_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single debtor by ID
   * @param {number} id - Debtor ID
   * @returns {Promise<Object>} - Debtor details
   */
  getById(id) {
    return apiClient.get(`${BASE_PATH}/${id}`)
  },

  /**
   * Create a new debtor
   * @param {Object} data - Debtor data
   * @param {string} data.name - Debtor name (required)
   * @param {string} [data.phone] - Phone number
   * @param {number} [data.amount] - Debt amount
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created debtor
   */
  create(data) {
    return apiClient.post(BASE_PATH, data)
  },

  /**
   * Update an existing debtor
   * @param {number} id - Debtor ID
   * @param {Object} data - Updated debtor data
   * @returns {Promise<Object>} - Updated debtor
   */
  update(id, data) {
    return apiClient.put(`${BASE_PATH}/${id}`, data)
  },

  /**
   * Delete a debtor
   * @param {number} id - Debtor ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${BASE_PATH}/${id}`)
  },
}

export default debtorsApi
