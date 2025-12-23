import { apiFetch, setTokens, clearTokens, getRefreshToken } from './client'

export async function login(credentials) {
  const result = await apiFetch('/auth/login', {
    method: 'POST',
    body: credentials,
    auth: false,
  })

  const payload = result?.data || {}
  setTokens(payload)
  return payload
}

export async function refresh() {
  const refreshToken = getRefreshToken()
  if (!refreshToken) throw new Error('Missing refresh token')

  const result = await apiFetch('/auth/refresh', {
    method: 'POST',
    body: { refreshToken },
    auth: false,
  })

  const payload = result?.data || {}
  setTokens(payload)
  return payload
}

export async function logout() {
  const refreshToken = getRefreshToken()
  await apiFetch('/auth/logout', {
    method: 'POST',
    body: { refreshToken },
    auth: false,
  })
  clearTokens()
}

export async function getMyPages() {
  const result = await apiFetch('/auth/pages', { method: 'GET', auth: true })
  return result?.data || []
}

