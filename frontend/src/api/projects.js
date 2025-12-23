import { apiFetch } from './client'

export async function listProjects(params = {}) {
  const query = new URLSearchParams(params).toString()
  const suffix = query ? `?${query}` : ''
  const result = await apiFetch(`/projects${suffix}`, { method: 'GET' })
  return result?.data || []
}

export async function getProject(id) {
  const result = await apiFetch(`/projects/${id}`, { method: 'GET' })
  return result?.data
}

export async function getProjectStats(params = {}) {
  const query = new URLSearchParams(params).toString()
  const suffix = query ? `?${query}` : ''
  const result = await apiFetch(`/projects/stats${suffix}`, { method: 'GET' })
  return result?.data || {}
}

export async function getProjectWorkdays(id) {
  const result = await apiFetch(`/projects/${id}/workdays`, { method: 'GET' })
  return result?.data || []
}

export async function createProject(payload) {
  const result = await apiFetch('/projects', {
    method: 'POST',
    body: payload,
  })
  return result?.data || result
}

export async function updateProject(id, payload) {
  const result = await apiFetch(`/projects/${id}`, {
    method: 'PUT',
    body: payload,
  })
  return result?.data || result
}

