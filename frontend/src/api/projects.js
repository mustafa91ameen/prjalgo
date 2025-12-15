import apiClient from './client'

const BASE_PATH = '/api/v1/projects'

export const projectsApi = {
  /**
   * Get all projects with pagination
   * @param {Object} params - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @returns {Promise<Object>} - Projects list with pagination info
   */
  getAll(params = {}) {
    const query = new URLSearchParams()
    if (params.page) query.append('page', params.page)
    if (params.limit) query.append('limit', params.limit)
    const queryString = query.toString()
    return apiClient.get(`${BASE_PATH}${queryString ? `?${queryString}` : ''}`)
  },

  /**
   * Get a single project by ID
   * @param {number} id - Project ID
   * @returns {Promise<Object>} - Project details
   */
  getById(id) {
    return apiClient.get(`${BASE_PATH}/${id}`)
  },

  /**
   * Create a new project
   * @param {Object} data - Project data
   * @param {string} data.name - Project name (required)
   * @param {string} [data.type] - Project type
   * @param {string} [data.description] - Project description
   * @param {string} [data.clientPhone] - Client phone number
   * @param {string} [data.location] - Project location
   * @param {string} [data.startDate] - Start date (ISO format)
   * @param {number} [data.duration] - Duration in days
   * @param {number} data.warningCost - Warning cost threshold (required)
   * @param {number} data.totalCost - Total budget (required)
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Created project
   */
  create(data) {
    return apiClient.post(BASE_PATH, data)
  },

  /**
   * Update an existing project
   * @param {number} id - Project ID
   * @param {Object} data - Updated project data
   * @param {string} [data.name] - Project name
   * @param {string} [data.type] - Project type
   * @param {string} [data.description] - Project description
   * @param {string} [data.clientPhone] - Client phone number
   * @param {string} [data.location] - Project location
   * @param {string} [data.startDate] - Start date (ISO format)
   * @param {number} [data.duration] - Duration in days
   * @param {number} [data.warningCost] - Warning cost threshold
   * @param {number} [data.totalCost] - Total budget
   * @param {string} [data.status] - Project status
   * @param {number} [data.progressPercentage] - Progress (0-100)
   * @param {string} [data.notes] - Additional notes
   * @returns {Promise<Object>} - Updated project
   */
  update(id, data) {
    return apiClient.put(`${BASE_PATH}/${id}`, data)
  },

  /**
   * Delete a project
   * @param {number} id - Project ID
   * @returns {Promise<null>}
   */
  delete(id) {
    return apiClient.delete(`${BASE_PATH}/${id}`)
  },

  /**
   * Get workdays for a specific project
   * @param {number} projectId - Project ID
   * @param {Object} [params] - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @returns {Promise<Object>} - Workdays list
   */
  getWorkdays(projectId, params = {}) {
    const query = new URLSearchParams()
    if (params.page) query.append('page', params.page)
    if (params.limit) query.append('limit', params.limit)
    const queryString = query.toString()
    return apiClient.get(`${BASE_PATH}/${projectId}/workdays${queryString ? `?${queryString}` : ''}`)
  },

  /**
   * Get expenses for a specific project
   * @param {number} projectId - Project ID
   * @param {Object} [params] - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @returns {Promise<Object>} - Expenses list
   */
  getExpenses(projectId, params = {}) {
    const query = new URLSearchParams()
    if (params.page) query.append('page', params.page)
    if (params.limit) query.append('limit', params.limit)
    const queryString = query.toString()
    return apiClient.get(`${BASE_PATH}/${projectId}/expenses${queryString ? `?${queryString}` : ''}`)
  },

  /**
   * Get team members for a specific project
   * @param {number} projectId - Project ID
   * @param {Object} [params] - Query parameters
   * @param {number} [params.page=1] - Page number
   * @param {number} [params.limit=10] - Items per page
   * @returns {Promise<Object>} - Team members list
   */
  getTeamMembers(projectId, params = {}) {
    const query = new URLSearchParams()
    if (params.page) query.append('page', params.page)
    if (params.limit) query.append('limit', params.limit)
    const queryString = query.toString()
    return apiClient.get(`${BASE_PATH}/${projectId}/team-members${queryString ? `?${queryString}` : ''}`)
  },
}

export default projectsApi
