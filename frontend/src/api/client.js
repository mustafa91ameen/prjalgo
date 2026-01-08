// API base URL - defaults to '/api/v1' for Docker (nginx proxy) or uses VITE_API_URL for local dev
const API_BASE_URL = import.meta.env.VITE_API_URL || '/api/v1'

const defaultHeaders = {
  'Content-Type': 'application/json',
  Accept: 'application/json',
}

function getAccessToken() {
  return localStorage.getItem('accessToken') || ''
}

export function getRefreshToken() {
  return localStorage.getItem('refreshToken') || ''
}

export function setTokens({ accessToken, refreshToken }) {
  if (accessToken) localStorage.setItem('accessToken', accessToken)
  if (refreshToken) localStorage.setItem('refreshToken', refreshToken)
}

export function clearTokens() {
  localStorage.removeItem('accessToken')
  localStorage.removeItem('refreshToken')
  localStorage.removeItem('isAuthenticated')
  localStorage.removeItem('user')
  localStorage.removeItem('username')
  localStorage.removeItem('loginTime')
}

// Track if we're currently refreshing to prevent multiple refresh calls
let isRefreshing = false
let refreshPromise = null

async function refreshAccessToken() {
  const refreshToken = getRefreshToken()
  if (!refreshToken) {
    throw new Error('No refresh token available')
  }

  const response = await fetch(`${API_BASE_URL}/auth/refresh`, {
    method: 'POST',
    headers: defaultHeaders,
    body: JSON.stringify({ refreshToken }),
  })

  const data = await response.json().catch(() => ({}))

  if (!response.ok) {
    throw new Error(data?.error || 'Token refresh failed')
  }

  const payload = data?.data || {}
  setTokens(payload)
  return payload.accessToken
}

async function handleTokenRefresh() {
  // If already refreshing, wait for that promise
  if (isRefreshing && refreshPromise) {
    return refreshPromise
  }

  isRefreshing = true
  refreshPromise = refreshAccessToken()
    .finally(() => {
      isRefreshing = false
      refreshPromise = null
    })

  return refreshPromise
}

export async function apiFetch(path, { method = 'GET', body, headers = {}, auth = true, _isRetry = false } = {}) {
  const finalHeaders = { ...defaultHeaders, ...headers }

  if (auth) {
    const token = getAccessToken()
    if (token) {
      finalHeaders.Authorization = `Bearer ${token}`
    }
  }

  let response
  try {
    response = await fetch(`${API_BASE_URL}${path}`, {
      method,
      headers: finalHeaders,
      body: body ? JSON.stringify(body) : undefined,
    })
  } catch (err) {
    const error = new Error('فشل الاتصال بالخادم')
    error.status = 0
    throw error
  }

  // Handle 401 - try to refresh token and retry once
  if (response.status === 401 && auth && !_isRetry) {
    try {
      await handleTokenRefresh()
      // Retry the original request with new token
      return apiFetch(path, { method, body, headers, auth, _isRetry: true })
    } catch (refreshError) {
      // Refresh failed - clear tokens and redirect to login
      clearTokens()
      window.location.href = '/login'
      const error = new Error('انتهت صلاحية الجلسة، يرجى تسجيل الدخول مرة أخرى')
      error.status = 401
      throw error
    }
  }

  const data = await response.json().catch(() => ({}))

  if (!response.ok) {
    const message = data?.error || response.statusText || 'Request failed'
    const error = new Error(message)
    error.status = response.status
    error.data = data
    throw error
  }

  return data
}
