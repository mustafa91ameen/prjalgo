import { apiFetch } from './client'
import { DEFAULT_PAGE, DEFAULT_LIMIT } from '../constants/pagination'

// ============ Work Categories ============

export async function listCategories({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  return apiFetch(`/workCategories?${query}`, { method: 'GET' })
}

export async function getCategory(id) {
  const result = await apiFetch(`/workCategories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createCategory(payload) {
  return apiFetch(`/workCategories`, {
    method: 'POST',
    body: payload,
  })
}

export async function updateCategory(id, payload) {
  return apiFetch(`/workCategories/${id}`, {
    method: 'PUT',
    body: payload,
  })
}

export async function deleteCategory(id) {
  return apiFetch(`/workCategories/${id}`, {
    method: 'DELETE',
  })
}

export async function getCategoryStats({ period = '' } = {}) {
  const query = period ? `?period=${period}` : ''
  const result = await apiFetch(`/workCategories/stats${query}`, { method: 'GET' })
  return result?.data || result
}

// ============ Work SubCategories ============

export async function listSubCategories({ page = DEFAULT_PAGE, limit = DEFAULT_LIMIT } = {}) {
  const query = new URLSearchParams({ page, limit }).toString()
  return apiFetch(`/workSubcategories?${query}`, { method: 'GET' })
}

export async function getSubCategory(id) {
  const result = await apiFetch(`/workSubcategories/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function createSubCategory(payload) {
  return apiFetch(`/workSubcategories`, {
    method: 'POST',
    body: payload,
  })
}

export async function updateSubCategory(id, payload) {
  return apiFetch(`/workSubcategories/${id}`, {
    method: 'PUT',
    body: payload,
  })
}

export async function deleteSubCategory(id) {
  return apiFetch(`/workSubcategories/${id}`, {
    method: 'DELETE',
  })
}
