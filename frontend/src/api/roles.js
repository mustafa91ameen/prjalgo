import apiClient from './client'

const ROLES_PATH = '/api/v1/roles'
const PAGES_PATH = '/api/v1/pages'
const ROLE_PAGES_PATH = '/api/v1/role-pages'

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
  if (params.roleId) query.append('roleId', params.roleId)
  if (params.pageId) query.append('pageId', params.pageId)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

// ==================== ROLES ====================

export const rolesApi = {
  /**
   * Get all roles with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Roles list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${ROLES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single role by ID
   * @param {number} id - Role ID
   * @returns {Promise<Object>} - Role details
   */
  getById(id) {
    return apiClient.get(`${ROLES_PATH}/${id}`)
  },

  /**
   * Create a new role
   * @param {Object} data - Role data
   * @param {string} data.name - Role name (required)
   * @param {string} [data.description] - Role description
   * @returns {Promise<Object>} - Created role
   */
  create(data) {
    return apiClient.post(ROLES_PATH, data)
  },

  /**
   * Update an existing role
   * @param {number} id - Role ID
   * @param {Object} data - Updated role data
   * @returns {Promise<Object>} - Updated role
   */
  update(id, data) {
    return apiClient.put(`${ROLES_PATH}/${id}`, data)
  },

  /**
   * Delete a role
   * @param {number} id - Role ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${ROLES_PATH}/${id}`)
  },
}

// ==================== PAGES ====================

export const pagesApi = {
  /**
   * Get all pages with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Pages list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${PAGES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single page by ID
   * @param {number} id - Page ID
   * @returns {Promise<Object>} - Page details
   */
  getById(id) {
    return apiClient.get(`${PAGES_PATH}/${id}`)
  },

  /**
   * Create a new page
   * @param {Object} data - Page data
   * @param {string} data.name - Page name (required)
   * @param {string} data.route - Page route (required)
   * @param {string} [data.icon] - Page icon
   * @param {number} [data.order] - Display order
   * @returns {Promise<Object>} - Created page
   */
  create(data) {
    return apiClient.post(PAGES_PATH, data)
  },

  /**
   * Update an existing page
   * @param {number} id - Page ID
   * @param {Object} data - Updated page data
   * @returns {Promise<Object>} - Updated page
   */
  update(id, data) {
    return apiClient.put(`${PAGES_PATH}/${id}`, data)
  },

  /**
   * Delete a page
   * @param {number} id - Page ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${PAGES_PATH}/${id}`)
  },
}

// ==================== ROLE PAGES ====================

export const rolePagesApi = {
  /**
   * Get all role-page mappings with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {number} [params.roleId] - Filter by role
   * @param {number} [params.pageId] - Filter by page
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Role-page mappings list
   */
  getAll(params = {}) {
    return apiClient.get(`${ROLE_PAGES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single role-page mapping by ID
   * @param {number} id - Role-page mapping ID
   * @returns {Promise<Object>} - Role-page mapping details
   */
  getById(id) {
    return apiClient.get(`${ROLE_PAGES_PATH}/${id}`)
  },

  /**
   * Create a new role-page mapping (assign page to role with permissions)
   * @param {Object} data - Role-page data
   * @param {number} data.roleId - Role ID (required)
   * @param {number} data.pageId - Page ID (required)
   * @param {boolean} [data.canRead=false] - Read permission
   * @param {boolean} [data.canWrite=false] - Write permission
   * @param {boolean} [data.canDelete=false] - Delete permission
   * @returns {Promise<Object>} - Created role-page mapping
   */
  create(data) {
    return apiClient.post(ROLE_PAGES_PATH, data)
  },

  /**
   * Update an existing role-page mapping
   * @param {number} id - Role-page mapping ID
   * @param {Object} data - Updated permissions
   * @param {boolean} [data.canRead] - Read permission
   * @param {boolean} [data.canWrite] - Write permission
   * @param {boolean} [data.canDelete] - Delete permission
   * @returns {Promise<Object>} - Updated role-page mapping
   */
  update(id, data) {
    return apiClient.put(`${ROLE_PAGES_PATH}/${id}`, data)
  },

  /**
   * Delete a role-page mapping
   * @param {number} id - Role-page mapping ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${ROLE_PAGES_PATH}/${id}`)
  },
}

export default {
  roles: rolesApi,
  pages: pagesApi,
  rolePages: rolePagesApi,
}
