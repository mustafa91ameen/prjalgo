import apiClient from './client'

const BASE_PATH = '/api/v1/team-members'

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
  if (params.userId) query.append('userId', params.userId)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

export const teamsApi = {
  /**
   * Get all team members with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {number} [params.projectId] - Filter by project
   * @param {number} [params.userId] - Filter by user
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Team members list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${BASE_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single team member by ID
   * @param {number} id - Team member ID
   * @returns {Promise<Object>} - Team member details
   */
  getById(id) {
    return apiClient.get(`${BASE_PATH}/${id}`)
  },

  /**
   * Create a new team member (assign user to project)
   * @param {Object} data - Team member data
   * @param {number} data.userId - User ID (required)
   * @param {number} data.projectId - Project ID (required)
   * @param {string} [data.role] - Role in the project
   * @returns {Promise<Object>} - Created team member
   */
  create(data) {
    return apiClient.post(BASE_PATH, data)
  },

  /**
   * Delete a team member (remove user from project)
   * @param {number} id - Team member ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${BASE_PATH}/${id}`)
  },
}

export default teamsApi
