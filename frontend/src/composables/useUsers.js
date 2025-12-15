import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useUsersStore } from '@/stores/users'
import { usePermissions } from './usePermissions'

/**
 * Composable for users management
 * Wraps the users store with permission checks
 */
export function useUsers() {
  const usersStore = useUsersStore()
  const { canRead, canWrite, canDelete } = usePermissions()

  // Destructure reactive state from store
  const { users, currentUser, loading, error, pagination, filters } = storeToRefs(usersStore)

  // Permission checks
  const canReadUsers = canRead('/users')
  const canWriteUsers = canWrite('/users')
  const canDeleteUsers = canDelete('/users')

  // Computed statistics
  const totalUsers = computed(() => pagination.value.total || users.value.length)

  const activeUsers = computed(() => {
    return users.value.filter(u => u.status === 'active').length
  })

  const pendingUsers = computed(() => {
    return users.value.filter(u => u.status === 'pending').length
  })

  const inactiveUsers = computed(() => {
    return users.value.filter(u => u.status === 'inactive').length
  })

  const adminUsers = computed(() => {
    return users.value.filter(u => u.role === 'admin').length
  })

  // Filter users by role
  const getUsersByRole = (role) => {
    if (!role) return users.value
    return users.value.filter(u => u.role === role)
  }

  // Filter users by status
  const getUsersByStatus = (status) => {
    if (!status) return users.value
    return users.value.filter(u => u.status === status)
  }

  // Search users
  const searchUsers = (query) => {
    if (!query) return users.value
    const lowerQuery = query.toLowerCase()
    return users.value.filter(u =>
      u.name?.toLowerCase().includes(lowerQuery) ||
      u.email?.toLowerCase().includes(lowerQuery) ||
      u.phone?.includes(query) ||
      u.department?.toLowerCase().includes(lowerQuery)
    )
  }

  // Actions with permission checks
  const fetchUsers = async (params = {}) => {
    if (!canReadUsers.value) {
      console.warn('No permission to read users')
      return
    }
    return usersStore.fetchUsers(params)
  }

  const fetchUserById = async (id) => {
    if (!canReadUsers.value) {
      console.warn('No permission to read users')
      return
    }
    return usersStore.fetchUserById(id)
  }

  const createUser = async (data) => {
    if (!canWriteUsers.value) {
      console.warn('No permission to create users')
      return
    }
    return usersStore.createUser(data)
  }

  const updateUser = async (id, data) => {
    if (!canWriteUsers.value) {
      console.warn('No permission to update users')
      return
    }
    return usersStore.updateUser(id, data)
  }

  const deleteUser = async (id) => {
    if (!canDeleteUsers.value) {
      console.warn('No permission to delete users')
      return
    }
    return usersStore.deleteUser(id)
  }

  const changePassword = async (id, password) => {
    if (!canWriteUsers.value) {
      console.warn('No permission to change password')
      return
    }
    return usersStore.changePassword(id, password)
  }

  const updateStatus = async (id, status) => {
    if (!canWriteUsers.value) {
      console.warn('No permission to update user status')
      return
    }
    return usersStore.updateStatus(id, status)
  }

  // Filter and pagination helpers
  const setFilters = (newFilters) => {
    usersStore.setFilters(newFilters)
  }

  const clearFilters = () => {
    usersStore.clearFilters()
  }

  const setPage = (page) => {
    usersStore.setPage(page)
  }

  const setLimit = (limit) => {
    usersStore.setLimit(limit)
  }

  return {
    // State
    users,
    currentUser,
    loading,
    error,
    pagination,
    filters,

    // Computed statistics
    totalUsers,
    activeUsers,
    pendingUsers,
    inactiveUsers,
    adminUsers,

    // Permissions
    canReadUsers,
    canWriteUsers,
    canDeleteUsers,

    // Methods
    fetchUsers,
    fetchUserById,
    createUser,
    updateUser,
    deleteUser,
    changePassword,
    updateStatus,
    getUsersByRole,
    getUsersByStatus,
    searchUsers,
    setFilters,
    clearFilters,
    setPage,
    setLimit
  }
}
