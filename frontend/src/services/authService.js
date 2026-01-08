import { login as apiLogin, refresh as apiRefresh, logout as apiLogout } from '@/api/auth'
import { clearTokens } from '@/api/client'

const ACCESS_KEY = 'accessToken'
const USER_KEY = 'user'
const AUTH_FLAG = 'isAuthenticated'
const USERNAME_KEY = 'username'

function persistSession({ accessToken, refreshToken, username, userId }) {
  if (accessToken) localStorage.setItem(ACCESS_KEY, accessToken)
  if (refreshToken) localStorage.setItem('refreshToken', refreshToken)
  if (username) localStorage.setItem(USERNAME_KEY, username)
  if (userId) {
    localStorage.setItem(USER_KEY, JSON.stringify({ id: userId, username }))
  }
  localStorage.setItem(AUTH_FLAG, 'true')
  localStorage.setItem('loginTime', new Date().toISOString())
}

export function getCurrentUser() {
  const raw = localStorage.getItem(USER_KEY)
  if (!raw) return null
  try {
    return JSON.parse(raw)
  } catch (err) {
    console.warn('Failed to parse user from storage', err)
    return null
  }
}

export function isAuthenticated() {
  return localStorage.getItem(AUTH_FLAG) === 'true' && Boolean(localStorage.getItem(ACCESS_KEY))
}

export async function handleLogin(credentials) {
  const session = await apiLogin(credentials)
  persistSession(session)
  return session
}

export async function tryRefresh() {
  const session = await apiRefresh()
  persistSession(session)
  return session
}

export async function handleLogout() {
  try {
    await apiLogout()
  } finally {
    clearTokens()
  }
}
