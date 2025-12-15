import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import { usersApi } from '@/api/users'
import { teamsApi } from '@/api/teams'
import { projectsApi } from '@/api/projects'

export const useEngineersStore = defineStore('engineers', () => {
  // State
  const engineers = ref([])
  const currentEngineer = ref(null)
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
  const availableProjects = ref([])

  // Getters
  const engineerCount = computed(() => engineers.value.length)
  const totalEngineers = computed(() => pagination.value.total)
  const hasEngineers = computed(() => engineers.value.length > 0)
  const isLoading = computed(() => loading.value)
  const activeEngineers = computed(() =>
    engineers.value.filter((e) => e.status === 'active')
  )
  const inactiveEngineers = computed(() =>
    engineers.value.filter((e) => e.status === 'inactive')
  )
  const averageRating = computed(() => {
    if (engineers.value.length === 0) return 0
    const total = engineers.value.reduce((sum, e) => sum + (e.rating || 0), 0)
    return (total / engineers.value.length).toFixed(1)
  })
  const totalProjects = computed(() => {
    return engineers.value.reduce((total, e) => total + (e.projects?.length || 0), 0)
  })

  /**
   * Fetch all engineers (users) with pagination
   * @param {Object} params - Query parameters
   */
  async function fetchEngineers(params = {}) {
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

      // Transform users to engineers format
      const usersData = response.data ?? []
      engineers.value = usersData.map(user => ({
        ...user,
        name: user.fullName || user.username,
        email: user.email || `${user.username}@example.com`,
        phone: user.phone || '',
        specialization: user.jobTitle || 'مهندس',
        rating: user.rating || 0,
        skills: user.skills || [],
        experience: user.experience || '',
        avatar: user.avatar || null,
        projects: user.projects || [],
      }))

      pagination.value = {
        page: response.page ?? 1,
        limit: response.limit ?? 10,
        total: response.total ?? 0,
        totalPages: response.totalPages ?? 0,
      }
    } catch (err) {
      error.value = err.message
      engineers.value = []
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch a single engineer by ID
   * @param {number} id - Engineer/User ID
   */
  async function fetchEngineerById(id) {
    loading.value = true
    error.value = null

    try {
      const response = await usersApi.getById(id)
      const user = response.data ?? response
      currentEngineer.value = {
        ...user,
        name: user.fullName || user.username,
        email: user.email || `${user.username}@example.com`,
        phone: user.phone || '',
        specialization: user.jobTitle || 'مهندس',
        rating: user.rating || 0,
        skills: user.skills || [],
        experience: user.experience || '',
        avatar: user.avatar || null,
        projects: user.projects || [],
      }
      return currentEngineer.value
    } catch (err) {
      error.value = err.message
      currentEngineer.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Create a new engineer (user)
   * @param {Object} data - Engineer data
   * @returns {Promise<Object|null>} - Created engineer or null
   */
  async function createEngineer(data) {
    loading.value = true
    error.value = null

    try {
      const userData = {
        username: data.name.toLowerCase().replace(/\s+/g, '_'),
        password: data.password,
        fullName: data.name,
        email: data.email || `${data.name.toLowerCase().replace(/\s+/g, '')}@example.com`,
        phone: data.phone || '',
        jobTitle: data.specialization || 'مهندس',
        status: 'active',
      }

      const response = await usersApi.create(userData)
      const newUser = response.data ?? response
      const newEngineer = {
        ...newUser,
        name: newUser.fullName || newUser.username,
        email: newUser.email,
        phone: newUser.phone || '',
        specialization: newUser.jobTitle || 'مهندس',
        rating: 0,
        skills: [],
        experience: '',
        avatar: null,
        projects: [],
      }

      engineers.value.unshift(newEngineer)
      pagination.value.total += 1
      toast.success('تم إضافة المهندس بنجاح')
      return newEngineer
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Update an existing engineer
   * @param {number} id - Engineer ID
   * @param {Object} data - Updated data
   * @returns {Promise<Object|null>} - Updated engineer or null
   */
  async function updateEngineer(id, data) {
    loading.value = true
    error.value = null

    try {
      const userData = {
        fullName: data.name,
        email: data.email,
        phone: data.phone,
        jobTitle: data.specialization,
      }

      // If password is provided, update it separately
      if (data.password) {
        await usersApi.changePassword(id, { password: data.password })
      }

      const response = await usersApi.update(id, userData)
      const updatedUser = response.data ?? response
      const updatedEngineer = {
        ...updatedUser,
        name: updatedUser.fullName || updatedUser.username,
        email: updatedUser.email,
        phone: updatedUser.phone || '',
        specialization: updatedUser.jobTitle || 'مهندس',
        rating: data.rating || 0,
        skills: data.skills || [],
        experience: data.experience || '',
        avatar: data.avatar || null,
        projects: data.projects || [],
      }

      const index = engineers.value.findIndex((e) => e.id === id)
      if (index !== -1) {
        engineers.value[index] = updatedEngineer
      }

      if (currentEngineer.value?.id === id) {
        currentEngineer.value = updatedEngineer
      }

      toast.success('تم تحديث المهندس بنجاح')
      return updatedEngineer
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Delete an engineer
   * @param {number} id - Engineer ID
   * @returns {Promise<boolean>} - Success status
   */
  async function deleteEngineer(id) {
    loading.value = true
    error.value = null

    try {
      await usersApi.delete(id)

      engineers.value = engineers.value.filter((e) => e.id !== id)
      pagination.value.total -= 1

      if (currentEngineer.value?.id === id) {
        currentEngineer.value = null
      }

      toast.success('تم حذف المهندس بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch available projects for assignment
   */
  async function fetchAvailableProjects() {
    try {
      const response = await projectsApi.getAll({ limit: 100 })
      availableProjects.value = (response.data ?? []).map(p => ({
        id: p.id,
        name: p.name,
        status: p.status,
      }))
    } catch {
      availableProjects.value = []
    }
  }

  /**
   * Add project to engineer (create team member)
   * @param {number} engineerId - Engineer ID
   * @param {number} projectId - Project ID
   */
  async function addProjectToEngineer(engineerId, projectId) {
    try {
      await teamsApi.create({
        userId: engineerId,
        projectId: projectId,
      })

      // Update local state
      const engineer = engineers.value.find(e => e.id === engineerId)
      if (engineer) {
        const project = availableProjects.value.find(p => p.id === projectId)
        if (project && !engineer.projects) {
          engineer.projects = []
        }
        if (project && !engineer.projects.some(p => p.id === projectId)) {
          engineer.projects.push(project)
        }
      }

      toast.success('تم إضافة المشروع بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    }
  }

  /**
   * Remove project from engineer (delete team member)
   * @param {number} engineerId - Engineer ID
   * @param {number} projectId - Project ID
   */
  async function removeProjectFromEngineer(engineerId, projectId) {
    try {
      // Need to find the team member ID first
      const response = await usersApi.getTeamMembers(engineerId)
      const teamMembers = response.data ?? []
      const teamMember = teamMembers.find(tm => tm.projectId === projectId)

      if (teamMember) {
        await teamsApi.delete(teamMember.id)
      }

      // Update local state
      const engineer = engineers.value.find(e => e.id === engineerId)
      if (engineer && engineer.projects) {
        engineer.projects = engineer.projects.filter(p => p.id !== projectId)
      }

      toast.success('تم إزالة المشروع بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    }
  }

  /**
   * Fetch engineer's projects (team members)
   * @param {number} engineerId - Engineer ID
   */
  async function fetchEngineerProjects(engineerId) {
    try {
      const response = await usersApi.getTeamMembers(engineerId)
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
   * Clear current engineer
   */
  function clearCurrentEngineer() {
    currentEngineer.value = null
  }

  /**
   * Reset store state
   */
  function $reset() {
    engineers.value = []
    currentEngineer.value = null
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
    availableProjects.value = []
  }

  return {
    // State
    engineers,
    currentEngineer,
    loading,
    error,
    pagination,
    filters,
    availableProjects,
    // Getters
    engineerCount,
    totalEngineers,
    hasEngineers,
    isLoading,
    activeEngineers,
    inactiveEngineers,
    averageRating,
    totalProjects,
    // Actions
    fetchEngineers,
    fetchEngineerById,
    createEngineer,
    updateEngineer,
    deleteEngineer,
    fetchAvailableProjects,
    addProjectToEngineer,
    removeProjectFromEngineer,
    fetchEngineerProjects,
    setFilters,
    clearFilters,
    setPage,
    setLimit,
    clearCurrentEngineer,
    $reset,
  }
})
