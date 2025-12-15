import apiClient from './client'

const BASE_PATH = '/api/v1/users'
const USER_ROLES_PATH = '/api/v1/user-roles'

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
  if (params.status) query.append('status', params.status)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

export const usersApi = {
  /**
   * Get all users with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {string} [params.status] - Filter by status
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Users list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${BASE_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single user by ID
   * @param {number} id - User ID
   * @returns {Promise<Object>} - User details
   */
  getById(id) {
    return apiClient.get(`${BASE_PATH}/${id}`)
  },

  /**
   * Create a new user
   * @param {Object} data - User data
   * @param {string} data.username - Username (required)
   * @param {string} data.password - Password (required)
   * @param {string} [data.fullName] - Full name
   * @param {string} [data.email] - Email address
   * @param {string} [data.phone] - Phone number
   * @returns {Promise<Object>} - Created user
   */
  create(data) {
    return apiClient.post(BASE_PATH, data)
  },

  /**
   * Update an existing user
   * @param {number} id - User ID
   * @param {Object} data - Updated user data
   * @returns {Promise<Object>} - Updated user
   */
  update(id, data) {
    return apiClient.put(`${BASE_PATH}/${id}`, data)
  },

  /**
   * Delete a user
   * @param {number} id - User ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${BASE_PATH}/${id}`)
  },

  /**
   * Change user password
   * @param {number} id - User ID
   * @param {Object} data - Password data
   * @param {string} data.password - New password
   * @returns {Promise<Object>} - Updated user
   */
  changePassword(id, data) {
    return apiClient.put(`${BASE_PATH}/${id}/password`, data)
  },

  /**
   * Update user status (active/inactive)
   * @param {number} id - User ID
   * @param {Object} data - Status data
   * @param {string} data.status - New status
   * @returns {Promise<Object>} - Updated user
   */
  updateStatus(id, data) {
    return apiClient.patch(`${BASE_PATH}/${id}/status`, data)
  },

  /**
   * Get user's roles
   * @param {number} id - User ID
   * @returns {Promise<Object>} - User's roles
   */
  getRoles(id) {
    return apiClient.get(`${BASE_PATH}/${id}/roles`)
  },

  /**
   * Get user's team members
   * @param {number} id - User ID
   * @param {Object} [params] - Query parameters
   * @returns {Promise<Object>} - User's team members
   */
  getTeamMembers(id, params = {}) {
    return apiClient.get(`${BASE_PATH}/${id}/team-members${buildQuery(params)}`)
  },
}

export const userRolesApi = {
  /**
   * Get all user-role assignments
   * @param {Object} params - Query parameters
   * @returns {Promise<Object>} - User roles list
   */
  getAll(params = {}) {
    return apiClient.get(`${USER_ROLES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single user-role assignment by ID
   * @param {number} id - User role ID
   * @returns {Promise<Object>} - User role details
   */
  getById(id) {
    return apiClient.get(`${USER_ROLES_PATH}/${id}`)
  },

  /**
   * Assign a role to a user
   * @param {Object} data - Assignment data
   * @param {number} data.userId - User ID
   * @param {number} data.roleId - Role ID
   * @returns {Promise<Object>} - Created assignment
   */
  assign(data) {
    return apiClient.post(USER_ROLES_PATH, data)
  },

  /**
   * Remove a role from a user
   * @param {number} id - User role assignment ID
   * @returns {Promise<null>}
   */
  remove(id) {
    return apiClient.delete(`${USER_ROLES_PATH}/${id}`)
  },
}

export default usersApi
