import { apiFetch } from './client'

// ============ Income ============

export async function listIncome() {
  const result = await apiFetch('/income', { method: 'GET' })
  // Handle paginated response: {success: true, data: {data: [...], total, page, limit, totalPages}}
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function getIncomeById(id) {
  const result = await apiFetch(`/income/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getIncomeStats() {
  const result = await apiFetch('/income/stats', { method: 'GET' })
  return result?.data || result
}

export async function createIncome(payload) {
  const result = await apiFetch('/income', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateIncome(id, payload) {
  const result = await apiFetch(`/income/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteIncome(id) {
  const result = await apiFetch(`/income/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
