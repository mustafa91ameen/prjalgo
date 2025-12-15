import apiClient from './client'

const BASE_PATH = '/api/v1/auth'

export const authApi = {
  /**
   * Login user with username and password
   * @param {Object} credentials - Login credentials
   * @param {string} credentials.username - Username
   * @param {string} credentials.password - Password
   * @returns {Promise<Object>} - Access token, refresh token, user info
   */
  login(credentials) {
    return apiClient.post(`${BASE_PATH}/login`, credentials)
  },

  /**
   * Logout user
   * @param {string} refreshToken - Refresh token to invalidate
   * @returns {Promise<null>}
   */
  logout(refreshToken) {
    return apiClient.post(`${BASE_PATH}/logout`, { refreshToken })
  },

  /**
   * Refresh access token
   * @param {string} refreshToken - Current refresh token
   * @returns {Promise<Object>} - New access token, refresh token, user info
   */
  refresh(refreshToken) {
    return apiClient.post(`${BASE_PATH}/refresh`, { refreshToken })
  },

  /**
   * Get authenticated user's accessible pages with permissions
   * @returns {Promise<Object>} - List of pages with permissions
   */
  getPages() {
    return apiClient.get(`${BASE_PATH}/pages`)
  },
}

export default authApi
