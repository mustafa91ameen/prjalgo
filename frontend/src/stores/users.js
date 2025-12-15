import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import { usersApi } from '@/api/users'

export const useUsersStore = defineStore('users', () => {
  // State
  const users = ref([])
  const currentUser = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0,
    totalPages: 0,
  })
  const filters = ref({
    status: null,
    search: '',
  })

  // Getters
  const userCount = computed(() => users.value.length)
  const totalUsers = computed(() => pagination.value.total)
  const hasUsers = computed(() => users.value.length > 0)
  const isLoading = computed(() => loading.value)
  const activeUsers = computed(() => users.value.filter((u) => u.status === 'active'))
  const inactiveUsers = computed(() => users.value.filter((u) => u.status === 'inactive'))

  /**
   * Fetch all users with pagination
   * @param {Object} params - Query parameters
   */
  async function fetchUsers(params = {}) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.getAll({
        page: params.page || pagination.value.page,
        limit: params.limit || pagination.value.limit,
        status: params.status ?? filters.value.status,
        search: params.search ?? filters.value.search,
        ...params,
      })

      users.value = response.data ?? []
      pagination.value = {
        page: response.page ?? 1,
        limit: response.limit ?? 10,
        total: response.total ?? 0,
        totalPages: response.totalPages ?? 0,
      }
    } catch (err) {
      error.value = err.message
      users.value = []
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch a single user by ID
   * @param {number} id - User ID
   */
  async function fetchUserById(id) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.getById(id)
      currentUser.value = response.data ?? response
      return currentUser.value
    } catch (err) {
      error.value = err.message
      currentUser.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Create a new user
   * @param {Object} data - User data
   * @returns {Promise<Object|null>} - Created user or null
   */
  async function createUser(data) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.create(data)
      const newUser = response.data ?? response
      users.value.unshift(newUser)
      pagination.value.total += 1
      toast.success('تم إنشاء المستخدم بنجاح')
      return newUser
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Update an existing user
   * @param {number} id - User ID
   * @param {Object} data - Updated user data
   * @returns {Promise<Object|null>} - Updated user or null
   */
  async function updateUser(id, data) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.update(id, data)
      const updatedUser = response.data ?? response

      const index = users.value.findIndex((u) => u.id === id)
      if (index !== -1) {
        users.value[index] = updatedUser
      }

      if (currentUser.value?.id === id) {
        currentUser.value = updatedUser
      }

      toast.success('تم تحديث المستخدم بنجاح')
      return updatedUser
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Delete a user
   * @param {number} id - User ID
   * @returns {Promise<boolean>} - Success status
   */
  async function deleteUser(id) {
    loading.value = true
    error.value = null

    try {
      await usersApi.delete(id)

      users.value = users.value.filter((u) => u.id !== id)
      pagination.value.total -= 1

      if (currentUser.value?.id === id) {
        currentUser.value = null
      }

      toast.success('تم حذف المستخدم بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Change user password
   * @param {number} id - User ID
   * @param {string} password - New password
   * @returns {Promise<boolean>} - Success status
   */
  async function changePassword(id, password) {
    loading.value = true
    error.value = null

    try {
      await usersApi.changePassword(id, { password })
      toast.success('تم تغيير كلمة المرور بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Update user status
   * @param {number} id - User ID
   * @param {string} status - New status
   * @returns {Promise<boolean>} - Success status
   */
  async function updateStatus(id, status) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.updateStatus(id, { status })
      const updatedUser = response.data ?? response

      const index = users.value.findIndex((u) => u.id === id)
      if (index !== -1) {
        users.value[index] = updatedUser
      }

      toast.success('تم تحديث حالة المستخدم بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Get user roles
   * @param {number} id - User ID
   * @returns {Promise<Array>} - User roles
   */
  async function fetchUserRoles(id) {
    try {
      const response = await usersApi.getRoles(id)
      return response.data ?? []
    } catch {
      return []
    }
  }

  /**
   * Get user team members
   * @param {number} id - User ID
   * @param {Object} params - Query parameters
   * @returns {Promise<Array>} - Team members
   */
  async function fetchUserTeamMembers(id, params = {}) {
    try {
      const response = await usersApi.getTeamMembers(id, params)
      return response.data ?? []
    } catch {
      return []
    }
  }

  /**
   * Set filters
   * @param {Object} newFilters - Filter values
   */
  function setFilters(newFilters) {
    filters.value = { ...filters.value, ...newFilters }
    pagination.value.page = 1
  }

  /**
   * Clear filters
   */
  function clearFilters() {
    filters.value = {
      status: null,
      search: '',
    }
    pagination.value.page = 1
  }

  /**
   * Set current page for pagination
   * @param {number} page - Page number
   */
  function setPage(page) {
    pagination.value.page = page
  }

  /**
   * Set items per page
   * @param {number} limit - Items per page
   */
  function setLimit(limit) {
    pagination.value.limit = limit
    pagination.value.page = 1
  }

  /**
   * Clear current user
   */
  function clearCurrentUser() {
    currentUser.value = null
  }

  /**
   * Reset store state
   */
  function $reset() {
    users.value = []
    currentUser.value = null
    loading.value = false
    error.value = null
    pagination.value = {
      page: 1,
      limit: 10,
      total: 0,
      totalPages: 0,
    }
    filters.value = {
      status: null,
      search: '',
    }
  }

  return {
    // State
    users,
    currentUser,
    loading,
    error,
    pagination,
    filters,
    // Getters
    userCount,
    totalUsers,
    hasUsers,
    isLoading,
    activeUsers,
    inactiveUsers,
    // Actions
    fetchUsers,
    fetchUserById,
    createUser,
    updateUser,
    deleteUser,
    changePassword,
    updateStatus,
    fetchUserRoles,
    fetchUserTeamMembers,
    setFilters,
    clearFilters,
    setPage,
    setLimit,
    clearCurrentUser,
    $reset,
  }
})
