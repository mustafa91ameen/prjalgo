import { apiFetch } from './client'

// ============ Debtors ============

export async function listDebtors() {
  const result = await apiFetch('/debtors', { method: 'GET' })
  // Handle paginated response: {success: true, data: {data: [...], total, page, limit, totalPages}}
  if (Array.isArray(result?.data?.data)) return result.data.data
  if (Array.isArray(result?.data)) return result.data
  if (Array.isArray(result)) return result
  return []
}

export async function getDebtorById(id) {
  const result = await apiFetch(`/debtors/${id}`, { method: 'GET' })
  return result?.data || result
}

export async function getDebtorStats() {
  const result = await apiFetch('/debtors/stats', { method: 'GET' })
  return result?.data || result
}

export async function createDebtor(payload) {
  const result = await apiFetch('/debtors', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateDebtor(id, payload) {
  const result = await apiFetch(`/debtors/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

export async function deleteDebtor(id) {
  const result = await apiFetch(`/debtors/${id}`, {
    method: 'DELETE',
  })
  return result?.data || result
}
