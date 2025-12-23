import { apiFetch } from './client'

// ============ Work Categories ============

export async function listCategories({ page = 1, limit = 100 } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/work-categories?${query}`, { method: 'GET' })
  return result
}

export async function getCategory(id) {
  const result = await apiFetch(`/work-categories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createCategory(payload) {
  const result = await apiFetch(`/work-categories`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateCategory(id, payload) {
  const result = await apiFetch(`/work-categories/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteCategory(id) {
  const result = await apiFetch(`/work-categories/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

export async function getCategoryStats({ period = '' } = {}) {
  const query = period ? `?period=${period}` : ''
  const result = await apiFetch(`/work-categories/stats${query}`, { method: 'GET' })
  return result?.data || result
}

// ============ Work SubCategories ============

export async function listSubCategories({ page = 1, limit = 100 } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  const result = await apiFetch(`/work-subcategories?${query}`, { method: 'GET' })
  // Response structure: { success: true, data: { data: [...], total, page, limit, totalPages } }
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function getSubCategory(id) {
  const result = await apiFetch(`/work-subcategories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createSubCategory(payload) {
  const result = await apiFetch(`/work-subcategories`, {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateSubCategory(id, payload) {
  const result = await apiFetch(`/work-subcategories/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteSubCategory(id) {
  const result = await apiFetch(`/work-subcategories/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
