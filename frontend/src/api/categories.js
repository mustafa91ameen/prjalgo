import apiClient from './client'

const CATEGORIES_PATH = '/api/v1/work-categories'
const SUBCATEGORIES_PATH = '/api/v1/work-subcategories'

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
  if (params.categoryId) query.append('categoryId', params.categoryId)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

// ==================== WORK CATEGORIES ====================

export const workCategoriesApi = {
  /**
   * Get all work categories with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Categories list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${CATEGORIES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single work category by ID
   * @param {number} id - Category ID
   * @returns {Promise<Object>} - Category details
   */
  getById(id) {
    return apiClient.get(`${CATEGORIES_PATH}/${id}`)
  },

  /**
   * Create a new work category
   * @param {Object} data - Category data
   * @param {string} data.name - Category name (required)
   * @param {string} [data.description] - Category description
   * @returns {Promise<Object>} - Created category
   */
  create(data) {
    return apiClient.post(CATEGORIES_PATH, data)
  },

  /**
   * Update an existing work category
   * @param {number} id - Category ID
   * @param {Object} data - Updated category data
   * @returns {Promise<Object>} - Updated category
   */
  update(id, data) {
    return apiClient.put(`${CATEGORIES_PATH}/${id}`, data)
  },

  /**
   * Delete a work category
   * @param {number} id - Category ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${CATEGORIES_PATH}/${id}`)
  },
}

// ==================== WORK SUBCATEGORIES ====================

export const workSubcategoriesApi = {
  /**
   * Get all work subcategories with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {string} [params.search] - Search term
   * @param {number} [params.categoryId] - Filter by parent category
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Subcategories list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${SUBCATEGORIES_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single work subcategory by ID
   * @param {number} id - Subcategory ID
   * @returns {Promise<Object>} - Subcategory details
   */
  getById(id) {
    return apiClient.get(`${SUBCATEGORIES_PATH}/${id}`)
  },

  /**
   * Create a new work subcategory
   * @param {Object} data - Subcategory data
   * @param {string} data.name - Subcategory name (required)
   * @param {number} data.categoryId - Parent category ID (required)
   * @param {string} [data.description] - Subcategory description
   * @returns {Promise<Object>} - Created subcategory
   */
  create(data) {
    return apiClient.post(SUBCATEGORIES_PATH, data)
  },

  /**
   * Update an existing work subcategory
   * @param {number} id - Subcategory ID
   * @param {Object} data - Updated subcategory data
   * @returns {Promise<Object>} - Updated subcategory
   */
  update(id, data) {
    return apiClient.put(`${SUBCATEGORIES_PATH}/${id}`, data)
  },

  /**
   * Delete a work subcategory
   * @param {number} id - Subcategory ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${SUBCATEGORIES_PATH}/${id}`)
  },
}

export default {
  categories: workCategoriesApi,
  subcategories: workSubcategoriesApi,
}
