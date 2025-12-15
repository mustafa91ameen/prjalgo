import { storeToRefs } from 'pinia'
import { useProjectsStore } from '@/stores/projects'
import { usePermissions } from './usePermissions'

/**
 * Composable for project-related logic
 * Wraps the projects store with permission checks and additional utilities
 *
 * @returns {Object} Project utilities
 */
export function useProjects() {
  const projectsStore = useProjectsStore()
  const { canRead, canWrite, canDelete } = usePermissions()

  // Get reactive state from store using storeToRefs
  const { projects, currentProject, loading, error, pagination, projectCount, totalProjects, hasProjects, isLoading } =
    storeToRefs(projectsStore)

  // Permission checks for projects page
  const canReadProjects = canRead('/projects')
  const canWriteProjects = canWrite('/projects')
  const canDeleteProjects = canDelete('/projects')

  /**
   * Fetch all projects (with permission check)
   * @param {Object} params - Query parameters
   */
  async function fetchProjects(params = {}) {
    if (!canReadProjects) {
      return
    }
    return projectsStore.fetchProjects(params)
  }

  /**
   * Fetch a single project by ID
   * @param {number} id - Project ID
   */
  async function fetchProjectById(id) {
    if (!canReadProjects) {
      return null
    }
    return projectsStore.fetchProjectById(id)
  }

  /**
   * Create a new project (with permission check)
   * @param {Object} data - Project data
   */
  async function createProject(data) {
    if (!canWriteProjects) {
      return null
    }
    return projectsStore.createProject(data)
  }

  /**
   * Update a project (with permission check)
   * @param {number} id - Project ID
   * @param {Object} data - Updated data
   */
  async function updateProject(id, data) {
    if (!canWriteProjects) {
      return null
    }
    return projectsStore.updateProject(id, data)
  }

  /**
   * Delete a project (with permission check)
   * @param {number} id - Project ID
   */
  async function deleteProject(id) {
    if (!canDeleteProjects) {
      return false
    }
    return projectsStore.deleteProject(id)
  }

  /**
   * Get project workdays
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectWorkdays(projectId, params = {}) {
    return projectsStore.fetchProjectWorkdays(projectId, params)
  }

  /**
   * Get project expenses
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectExpenses(projectId, params = {}) {
    return projectsStore.fetchProjectExpenses(projectId, params)
  }

  /**
   * Get project team members
   * @param {number} projectId - Project ID
   * @param {Object} params - Query parameters
   */
  async function fetchProjectTeamMembers(projectId, params = {}) {
    return projectsStore.fetchProjectTeamMembers(projectId, params)
  }

  /**
   * Set current page for pagination
   * @param {number} page - Page number
   */
  function setPage(page) {
    projectsStore.setPage(page)
  }

  /**
   * Set items per page
   * @param {number} limit - Items per page
   */
  function setLimit(limit) {
    projectsStore.setLimit(limit)
  }

  /**
   * Clear current project
   */
  function clearCurrentProject() {
    projectsStore.clearCurrentProject()
  }

  /**
   * Reset store state
   */
  function resetProjects() {
    projectsStore.$reset()
  }

  return {
    // State (reactive refs)
    projects,
    currentProject,
    loading,
    error,
    pagination,
    // Getters (reactive)
    projectCount,
    totalProjects,
    hasProjects,
    isLoading,
    // Permissions
    canReadProjects,
    canWriteProjects,
    canDeleteProjects,
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
    resetProjects,
  }
}

export default useProjects
