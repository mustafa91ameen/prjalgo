import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import apiClient from '@/api/client'
import { authApi } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref(null)
  const pages = ref([])
  const loading = ref(false)
  const error = ref(null)

  // Getters
  const isAuthenticated = computed(() => !!apiClient.getToken())
  const userId = computed(() => user.value?.userId ?? null)
  const username = computed(() => user.value?.username ?? null)

  /**
   * Check if user has permission for a specific page and action
   * @param {string} route - Page route (e.g., '/projects')
   * @param {string} permission - Permission type ('read', 'write', 'delete')
   * @returns {boolean}
   */
  function hasPermission(route, permission) {
    const page = pages.value.find((p) => p.route === route)
    if (!page) return false
    return page.permissions?.includes(permission) ?? false
  }

  /**
   * Check if user can read a page
   * @param {string} route - Page route
   * @returns {boolean}
   */
  function canRead(route) {
    return hasPermission(route, 'read')
  }

  /**
   * Check if user can write to a page
   * @param {string} route - Page route
   * @returns {boolean}
   */
  function canWrite(route) {
    return hasPermission(route, 'write')
  }

  /**
   * Check if user can delete on a page
   * @param {string} route - Page route
   * @returns {boolean}
   */
  function canDelete(route) {
    return hasPermission(route, 'delete')
  }

  /**
   * Login user with username and password
   * @param {string} username - Username
   * @param {string} password - Password
   * @returns {Promise<boolean>} - Success status
   */
  async function login(usernameInput, passwordInput) {
    loading.value = true
    error.value = null

    try {
      const response = await authApi.login({
        username: usernameInput,
        password: passwordInput,
      })

      // Store tokens
      apiClient.setTokens(response.data.accessToken, response.data.refreshToken)

      // Store user info
      user.value = {
        userId: response.data.userId,
        username: response.data.username,
      }

      // Fetch user pages after successful login
      await fetchPages()

      toast.success('Login successful')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Logout user and clear all auth state
   */
  async function logout() {
    loading.value = true

    try {
      const refreshToken = apiClient.getRefreshToken()
      if (refreshToken) {
        await authApi.logout(refreshToken)
      }
    } catch {
      // Ignore logout errors - we still want to clear local state
    } finally {
      // Clear tokens and state
      apiClient.clearTokens()
      user.value = null
      pages.value = []
      loading.value = false
      error.value = null
    }
  }

  /**
   * Refresh access token using refresh token
   * @returns {Promise<boolean>} - Success status
   */
  async function refreshToken() {
    const currentRefreshToken = apiClient.getRefreshToken()
    if (!currentRefreshToken) {
      return false
    }

    try {
      const response = await authApi.refresh(currentRefreshToken)

      // Update tokens
      apiClient.setTokens(response.data.accessToken, response.data.refreshToken)

      // Update user info
      user.value = {
        userId: response.data.userId,
        username: response.data.username,
      }

      return true
    } catch {
      // Refresh failed - logout user
      await logout()
      return false
    }
  }

  /**
   * Fetch user's accessible pages with permissions
   */
  async function fetchPages() {
    if (!isAuthenticated.value) {
      pages.value = []
      return
    }

    try {
      const response = await authApi.getPages()
      pages.value = response.data ?? []
    } catch {
      pages.value = []
    }
  }

  /**
   * Initialize auth state from stored tokens
   * Call this on app startup
   */
  async function initAuth() {
    if (!apiClient.getToken()) {
      return
    }

    loading.value = true

    try {
      // Try to fetch pages to verify token is still valid
      await fetchPages()

      // If we have pages, token is valid - try to get user info via refresh
      if (pages.value.length > 0) {
        await refreshToken()
      }
    } catch {
      // Token invalid - clear everything
      apiClient.clearTokens()
      user.value = null
      pages.value = []
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    user,
    pages,
    loading,
    error,
    // Getters
    isAuthenticated,
    userId,
    username,
    // Actions
    login,
    logout,
    refreshToken,
    fetchPages,
    initAuth,
    // Permission helpers
    hasPermission,
    canRead,
    canWrite,
    canDelete,
  }
})
