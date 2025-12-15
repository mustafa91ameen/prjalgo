import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import { projectsApi } from '@/api/projects'

export const useProjectsStore = defineStore('projects', () => {
  // State
  const projects = ref([])
  const currentProject = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0,
    totalPages: 0,
  })

  // Getters
  const projectCount = computed(() => projects.value.length)
  const totalProjects = computed(() => pagination.value.total)
  const hasProjects = computed(() => projects.value.length > 0)
  const isLoading = computed(() => loading.value)

  /**
   * Fetch all projects with pagination
   * @param {Object} params - Query parameters
   */
  async function fetchProjects(params = {}) {
    loading.value = true
    error.value = null

    try {
      const response = await projectsApi.getAll({
        page: params.page || pagination.value.page,
        limit: params.limit || pagination.value.limit,
        ...params,
      })

      projects.value = response.data ?? []
      pagination.value = {
        page: response.page ?? 1,
        limit: response.limit ?? 10,
        total: response.total ?? 0,
        totalPages: response.totalPages ?? 0,
      }
    } catch (err) {
      error.value = err.message
      projects.value = []
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch a single project by ID
   * @param {number} id - Project ID
   */
  async function fetchProjectById(id) {
    loading.value = true
    error.value = null

    try {
      const response = await projectsApi.getById(id)
      currentProject.value = response.data ?? response
      return currentProject.value
    } catch (err) {
      error.value = err.message
      currentProject.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Create a new project
   * @param {Object} data - Project data
   * @returns {Promise<Object|null>} - Created project or null
   */
  async function createProject(data) {
    loading.value = true
    error.value = null

    try {
      const response = await projectsApi.create(data)
      const newProject = response.data ?? response
      projects.value.unshift(newProject)
      pagination.value.total += 1
      toast.success('تم إنشاء المشروع بنجاح')
      return newProject
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Update an existing project
   * @param {number} id - Project ID
   * @param {Object} data - Updated project data
   * @returns {Promise<Object|null>} - Updated project or null
   */
  async function updateProject(id, data) {
    loading.value = true
    error.value = null

    try {
      const response = await projectsApi.update(id, data)
      const updatedProject = response.data ?? response

      const index = projects.value.findIndex((p) => p.id === id)
      if (index !== -1) {
        projects.value[index] = updatedProject
      }

      if (currentProject.value?.id === id) {
        currentProject.value = updatedProject
      }

      toast.success('تم تحديث المشروع بنجاح')
      return updatedProject
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Delete a project
   * @param {number} id - Project ID
   * @returns {Promise<boolean>} - Success status
   */
  async function deleteProject(id) {
    loading.value = true
    error.value = null

    try {
      await projectsApi.delete(id)

      projects.value = projects.value.filter((p) => p.id !== id)
      pagination.value.total -= 1

      if (currentProject.value?.id === id) {
        currentProject.value = null
      }

      toast.success('تم حذف المشروع بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
    }
  }

  /**
   * Get project workdays
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectWorkdays(projectId, params = {}) {
    try {
      const response = await projectsApi.getWorkdays(projectId, params)
      return response.data ?? []
    } catch {
      return []
    }
  }

  /**
   * Get project expenses
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectExpenses(projectId, params = {}) {
    try {
      const response = await projectsApi.getExpenses(projectId, params)
      return response.data ?? []
    } catch {
      return []
    }
  }

  /**
   * Get project team members
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectTeamMembers(projectId, params = {}) {
    try {
      const response = await projectsApi.getTeamMembers(projectId, params)
      return response.data ?? []
    } catch {
      return []
    }
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
   * Clear current project
   */
  function clearCurrentProject() {
    currentProject.value = null
  }

  /**
   * Reset store state
   */
  function $reset() {
    projects.value = []
    currentProject.value = null
    loading.value = false
    error.value = null
    pagination.value = {
      page: 1,
      limit: 10,
      total: 0,
      totalPages: 0,
    }
  }

  return {
    // State
    projects,
    currentProject,
    loading,
    error,
    pagination,
    // Getters
    projectCount,
    totalProjects,
    hasProjects,
    isLoading,
    // Actions
    fetchProjects,
    fetchProjectById,
    createProject,
    updateProject,
    deleteProject,
    fetchProjectWorkdays,
    fetchProjectExpenses,
    fetchProjectTeamMembers,
    setPage,
    setLimit,
    clearCurrentProject,
    $reset,
  }
})
