import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

export async function listWorkDays({ projectId, project_id, page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const projId = projectId || project_id
  if (projId) {
    const query = new URLSearchParams({ page, limit }).toString()
    const result = await apiFetch(`/projects/${projId}/workdays?${query}`, { method: 'GET' })
    const paginatedData = result?.data || {}
    if (Array.isArray(paginatedData)) {
      return { success: true, data: { items: paginatedData, total: paginatedData.length, page, limit, totalPages: 1 } }
    }
    return {
      success: true,
      data: {
        items: paginatedData.data || [],
        total: paginatedData.total || 0,
        page: paginatedData.page || page,
        limit: paginatedData.limit || limit,
        totalPages: paginatedData.totalPages || 0
      }
    }
  }
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/workdays?${query}`, { method: 'GET' })
  const paginatedData = result?.data || {}
  if (Array.isArray(paginatedData)) {
    return { success: true, data: { items: paginatedData, total: paginatedData.length, page, limit, totalPages: 1 } }
  }
  return {
    success: true,
    data: {
      items: paginatedData.data || [],
      total: paginatedData.total || 0,
      page: paginatedData.page || page,
      limit: paginatedData.limit || limit,
      totalPages: paginatedData.totalPages || 0
    }
  }
}

export async function getWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}`, { method: 'GET' })
  return { success: true, data: result?.data || result }
}

export async function createWorkDay(payload) {
  const result = await apiFetch(`/workdays`, {
    method: 'POST',
    body: payload,
  })
  return { success: true, data: result?.data || result }
}

export async function updateWorkDay(id, payload) {
  const result = await apiFetch(`/workdays/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return { success: true, data: result?.data || result }
}

export async function deleteWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}`, {
    method: 'DELETE',
  })
  return { success: true, data: result?.data || result }
}

export async function completeWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}/complete`, {
    method: 'PATCH',
  })
  return { success: true, data: result?.data || result }
}

export async function uncompleteWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}/uncomplete`, {
    method: 'PATCH',
  })
  return { success: true, data: result?.data || result }
}

export async function listWorkSubCategories({ page = 1, limit = 100 } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/workSubcategories?${query}`, { method: 'GET' })
  // Response structure: { success: true, data: { data: [...], total, page, limit, totalPages } }
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

// Aliases for consistency with imports
export { listWorkDays as listWorkdays }
export { getWorkDay as getWorkday }
export { createWorkDay as createWorkday }
export { updateWorkDay as updateWorkday }
export { deleteWorkDay as deleteWorkday }
export { completeWorkDay as completeWorkday }
export { uncompleteWorkDay as uncompleteWorkday }
