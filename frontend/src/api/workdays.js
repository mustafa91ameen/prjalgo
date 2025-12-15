import apiClient from './client'

const WORKDAYS_PATH = '/api/v1/workdays'
const MATERIALS_PATH = '/api/v1/workday-materials'
const LABOR_PATH = '/api/v1/workday-labor'
const EQUIPMENT_PATH = '/api/v1/workday-equipment'

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
  if (params.workdayId) query.append('workdayId', params.workdayId)
  if (params.startDate) query.append('startDate', params.startDate)
  if (params.endDate) query.append('endDate', params.endDate)
  if (params.status) query.append('status', params.status)
  if (params.sortBy) query.append('sortBy', params.sortBy)
  if (params.sortOrder) query.append('sortOrder', params.sortOrder)
  const queryString = query.toString()
  return queryString ? `?${queryString}` : ''
}

// ==================== WORKDAYS ====================

export const workdaysApi = {
  /**
   * Get all workdays with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {number} [params.projectId] - Filter by project
   * @param {string} [params.startDate] - Filter start date
   * @param {string} [params.endDate] - Filter end date
   * @param {string} [params.status] - Filter by status
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Workdays list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${WORKDAYS_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single workday by ID
   * @param {number} id - Workday ID
   * @returns {Promise<Object>} - Workday details
   */
  getById(id) {
    return apiClient.get(`${WORKDAYS_PATH}/${id}`)
  },

  /**
   * Create a new workday
   * @param {Object} data - Workday data
   * @param {number} data.projectId - Project ID (required)
   * @param {string} data.date - Workday date (required)
   * @param {string} [data.description] - Work description
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created workday
   */
  create(data) {
    return apiClient.post(WORKDAYS_PATH, data)
  },

  /**
   * Update an existing workday
   * @param {number} id - Workday ID
   * @param {Object} data - Updated workday data
   * @returns {Promise<Object>} - Updated workday
   */
  update(id, data) {
    return apiClient.put(`${WORKDAYS_PATH}/${id}`, data)
  },

  /**
   * Mark workday as complete
   * @param {number} id - Workday ID
   * @returns {Promise<Object>} - Updated workday
   */
  complete(id) {
    return apiClient.patch(`${WORKDAYS_PATH}/${id}/complete`, {})
  },

  /**
   * Delete a workday
   * @param {number} id - Workday ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${WORKDAYS_PATH}/${id}`)
  },
}

// ==================== MATERIALS ====================

export const workdayMaterialsApi = {
  /**
   * Get all workday materials with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {number} [params.workdayId] - Filter by workday
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Materials list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${MATERIALS_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single material by ID
   * @param {number} id - Material ID
   * @returns {Promise<Object>} - Material details
   */
  getById(id) {
    return apiClient.get(`${MATERIALS_PATH}/${id}`)
  },

  /**
   * Create a new material entry
   * @param {Object} data - Material data
   * @param {number} data.workdayId - Workday ID (required)
   * @param {string} data.name - Material name (required)
   * @param {number} data.quantity - Quantity (required)
   * @param {number} data.unitPrice - Unit price (required)
   * @param {string} [data.unit] - Unit of measurement
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created material
   */
  create(data) {
    return apiClient.post(MATERIALS_PATH, data)
  },

  /**
   * Update an existing material entry
   * @param {number} id - Material ID
   * @param {Object} data - Updated material data
   * @returns {Promise<Object>} - Updated material
   */
  update(id, data) {
    return apiClient.put(`${MATERIALS_PATH}/${id}`, data)
  },

  /**
   * Delete a material entry
   * @param {number} id - Material ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${MATERIALS_PATH}/${id}`)
  },
}

// ==================== LABOR ====================

export const workdayLaborApi = {
  /**
   * Get all labor entries with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {number} [params.workdayId] - Filter by workday
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Labor list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${LABOR_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single labor entry by ID
   * @param {number} id - Labor ID
   * @returns {Promise<Object>} - Labor details
   */
  getById(id) {
    return apiClient.get(`${LABOR_PATH}/${id}`)
  },

  /**
   * Create a new labor entry
   * @param {Object} data - Labor data
   * @param {number} data.workdayId - Workday ID (required)
   * @param {string} data.workerName - Worker name (required)
   * @param {number} data.hours - Hours worked (required)
   * @param {number} data.hourlyRate - Hourly rate (required)
   * @param {string} [data.role] - Worker role
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created labor entry
   */
  create(data) {
    return apiClient.post(LABOR_PATH, data)
  },

  /**
   * Update an existing labor entry
   * @param {number} id - Labor ID
   * @param {Object} data - Updated labor data
   * @returns {Promise<Object>} - Updated labor entry
   */
  update(id, data) {
    return apiClient.put(`${LABOR_PATH}/${id}`, data)
  },

  /**
   * Delete a labor entry
   * @param {number} id - Labor ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${LABOR_PATH}/${id}`)
  },
}

// ==================== EQUIPMENT ====================

export const workdayEquipmentApi = {
  /**
   * Get all equipment entries with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @param {number} [params.workdayId] - Filter by workday
   * @param {string} [params.sortBy] - Sort field
   * @param {string} [params.sortOrder] - Sort order (asc/desc)
   * @returns {Promise<Object>} - Equipment list with pagination info
   */
  getAll(params = {}) {
    return apiClient.get(`${EQUIPMENT_PATH}${buildQuery(params)}`)
  },

  /**
   * Get a single equipment entry by ID
   * @param {number} id - Equipment ID
   * @returns {Promise<Object>} - Equipment details
   */
  getById(id) {
    return apiClient.get(`${EQUIPMENT_PATH}/${id}`)
  },

  /**
   * Create a new equipment entry
   * @param {Object} data - Equipment data
   * @param {number} data.workdayId - Workday ID (required)
   * @param {string} data.name - Equipment name (required)
   * @param {number} data.quantity - Quantity (required)
   * @param {number} data.dailyRate - Daily rental rate (required)
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created equipment entry
   */
  create(data) {
    return apiClient.post(EQUIPMENT_PATH, data)
  },

  /**
   * Update an existing equipment entry
   * @param {number} id - Equipment ID
   * @param {Object} data - Updated equipment data
   * @returns {Promise<Object>} - Updated equipment entry
   */
  update(id, data) {
    return apiClient.put(`${EQUIPMENT_PATH}/${id}`, data)
  },

  /**
   * Delete an equipment entry
   * @param {number} id - Equipment ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${EQUIPMENT_PATH}/${id}`)
  },
}

export default {
  workdays: workdaysApi,
  materials: workdayMaterialsApi,
  labor: workdayLaborApi,
  equipment: workdayEquipmentApi,
}
