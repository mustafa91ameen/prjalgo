import { apiFetch } from './client'

/**
 * List all roles with pagination
 * @param {Object} params - Query parameters
 * @param {number} params.page - Page number
 * @param {number} params.limit - Items per page
 * @returns {Promise<Object>} Paginated response
 */
export async function listRoles(params = {}) {
  const query = new URLSearchParams()
  if (params.page) query.append('page', params.page.toString())
  if (params.limit) query.append('limit', params.limit.toString())
  const queryString = query.toString()
  return apiFetch(`/roles${queryString ? `?${queryString}` : '?limit=100'}`)
}

/**
 * Get a single role by ID
 * @param {number|string} id - Role ID
 * @returns {Promise<Object>} Role data
 */
export async function getRole(id) {
  const result = await apiFetch(`/roles/${id}`)
  return result?.data || result
}

/**
 * Create a new role
 * @param {Object} roleData - Role data
 * @param {string} roleData.name - Role name (required)
 * @param {string} roleData.description - Role description
 * @returns {Promise<Object>} Created role data
 */
export async function createRole(roleData) {
  const result = await apiFetch('/roles', {
    method: 'POST',
    body: roleData
  })
  return result?.data || result
}

/**
 * Update an existing role
 * @param {number|string} id - Role ID
 * @param {Object} roleData - Role data to update
 * @returns {Promise<Object>} Updated role data
 */
export async function updateRole(id, roleData) {
  const result = await apiFetch(`/roles/${id}`, {
    method: 'PUT',
    body: roleData
  })
  return result?.data || result
}

/**
 * Delete a role
 * @param {number|string} id - Role ID
 * @returns {Promise<Object>} Response
 */
export async function deleteRole(id) {
  const result = await apiFetch(`/roles/${id}`, {
    method: 'DELETE'
  })
  return result?.data || result
}

/**
 * Assign role to user
 * @param {number|string} userId - User ID
 * @param {number|string} roleId - Role ID
 * @returns {Promise<Object>} Response
 */
export async function assignRoleToUser(userId, roleId) {
  const result = await apiFetch('/userRoles', {
    method: 'POST',
    body: { userId, roleId }
  })
  return result?.data || result
}

/**
 * Remove role from user
 * @param {number|string} userRoleId - UserRole ID
 * @returns {Promise<Object>} Response
 */
export async function removeRoleFromUser(userRoleId) {
  const result = await apiFetch(`/userRoles/${userRoleId}`, {
    method: 'DELETE'
  })
  return result?.data || result
}

/**
 * Get role pages (permissions) by role ID
 * @param {number|string} roleId - Role ID
 * @returns {Promise<Object>} Response with role pages
 */
export async function getRolePages(roleId) {
  return apiFetch(`/rolePages/role/${roleId}`)
}
