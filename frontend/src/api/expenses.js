import { apiFetch } from './client'

// ============ Expenses ============

export async function listExpenses() {
  const result = await apiFetch('/expenses', { method: 'GET' })
  // Handle paginated response: {success: true, data: {data: [...], total, page, limit, totalPages}}
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function getExpenseById(id) {
  const result = await apiFetch(`/expenses/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getExpenseStats() {
  const result = await apiFetch('/expenses/stats', { method: 'GET' })
  return result?.data || result
}

export async function createExpense(payload) {
  const result = await apiFetch('/expenses', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateExpense(id, payload) {
  const result = await apiFetch(`/expenses/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteExpense(id) {
  const result = await apiFetch(`/expenses/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}

// ============ Project Expenses ============

export async function listExpensesByProject(projectId) {
  const result = await apiFetch(`/projects/${projectId}/expenses`, { method: 'GET' })
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}
