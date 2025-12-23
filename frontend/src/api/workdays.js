import { apiFetch } from './client'

export async function listWorkDays({ projectId, page = 1, limit = 20 } = {}) {
  if (projectId) {
    const result = await apiFetch(`/projects/${projectId}/workdays`, { method: 'GET' })
    return result?.data || result || []
  }
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/workdays?${query}`, { method: 'GET' })
  // Some endpoints may wrap data; normalize to array if possible
  if (Array.isArray(result)) return result
  if (Array.isArray(result?.data)) return result.data
  return []
}

export async function getWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createWorkDay(payload) {
  const result = await apiFetch(`/workdays`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateWorkDay(id, payload) {
  const result = await apiFetch(`/workdays/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteWorkDay(id) {
  const result = await apiFetch(`/workdays/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

export async function listWorkSubCategories({ page = 1, limit = 100 } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/work-subcategories?${query}`, { method: 'GET' })
  // Response structure: { success: true, data: { data: [...], total, page, limit, totalPages } }
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

