import { storeToRefs } from 'pinia'
import { useEngineersStore } from '@/stores/engineers'
import { usePermissions } from './usePermissions'

/**
 * Composable for engineer-related logic
 * Wraps the engineers store with permission checks and additional utilities
 *
 * @returns {Object} Engineer utilities
 */
export function useEngineers() {
  const engineersStore = useEngineersStore()
  const { canRead, canWrite, canDelete } = usePermissions()

  // Get reactive state from store using storeToRefs
  const {
    engineers,
    currentEngineer,
    loading,
    error,
    pagination,
    filters,
    availableProjects,
    engineerCount,
    totalEngineers,
    hasEngineers,
    isLoading,
    activeEngineers,
    inactiveEngineers,
    averageRating,
    totalProjects,
  } = storeToRefs(engineersStore)

  // Permission checks for engineers page
  const canReadEngineers = canRead('/engineers')
  const canWriteEngineers = canWrite('/engineers')
  const canDeleteEngineers = canDelete('/engineers')

  /**
   * Fetch all engineers (with permission check)
   * @param {Object} params - Query parameters
   */
  async function fetchEngineers(params = {}) {
    if (!canReadEngineers) {
      return
    }
    return engineersStore.fetchEngineers(params)
  }

  /**
   * Fetch a single engineer by ID
   * @param {number} id - Engineer ID
   */
  async function fetchEngineerById(id) {
    if (!canReadEngineers) {
      return null
    }
    return engineersStore.fetchEngineerById(id)
  }

  /**
   * Create a new engineer (with permission check)
   * @param {Object} data - Engineer data
   */
  async function createEngineer(data) {
    if (!canWriteEngineers) {
      return null
    }
    return engineersStore.createEngineer(data)
  }

  /**
   * Update an engineer (with permission check)
   * @param {number} id - Engineer ID
   * @param {Object} data - Updated data
   */
  async function updateEngineer(id, data) {
    if (!canWriteEngineers) {
      return null
    }
    return engineersStore.updateEngineer(id, data)
  }

  /**
   * Delete an engineer (with permission check)
   * @param {number} id - Engineer ID
   */
  async function deleteEngineer(id) {
    if (!canDeleteEngineers) {
      return false
    }
    return engineersStore.deleteEngineer(id)
  }

  /**
   * Fetch available projects for assignment
   */
  async function fetchAvailableProjects() {
    return engineersStore.fetchAvailableProjects()
  }

  /**
   * Add project to engineer
   * @param {number} engineerId - Engineer ID
   * @param {number} projectId - Project ID
   */
  async function addProjectToEngineer(engineerId, projectId) {
    if (!canWriteEngineers) {
      return false
    }
    return engineersStore.addProjectToEngineer(engineerId, projectId)
  }

  /**
   * Remove project from engineer
   * @param {number} engineerId - Engineer ID
   * @param {number} projectId - Project ID
   */
  async function removeProjectFromEngineer(engineerId, projectId) {
    if (!canWriteEngineers) {
      return false
    }
    return engineersStore.removeProjectFromEngineer(engineerId, projectId)
  }

  /**
   * Fetch engineer's projects
   * @param {number} engineerId - Engineer ID
   */
  async function fetchEngineerProjects(engineerId) {
    return engineersStore.fetchEngineerProjects(engineerId)
  }

  /**
   * Set filters
   * @param {Object} newFilters - Filter values
   */
  function setFilters(newFilters) {
    engineersStore.setFilters(newFilters)
  }

  /**
   * Clear filters
   */
  function clearFilters() {
    engineersStore.clearFilters()
  }

  /**
   * Set current page for pagination
   * @param {number} page - Page number
   */
  function setPage(page) {
    engineersStore.setPage(page)
  }

  /**
   * Set items per page
   * @param {number} limit - Items per page
   */
  function setLimit(limit) {
    engineersStore.setLimit(limit)
  }

  /**
   * Clear current engineer
   */
  function clearCurrentEngineer() {
    engineersStore.clearCurrentEngineer()
  }

  /**
   * Reset store state
   */
  function resetEngineers() {
    engineersStore.$reset()
  }

  return {
    // State (reactive refs)
    engineers,
    currentEngineer,
    loading,
    error,
    pagination,
    filters,
    availableProjects,
    // Getters (reactive)
    engineerCount,
    totalEngineers,
    hasEngineers,
    isLoading,
    activeEngineers,
    inactiveEngineers,
    averageRating,
    totalProjects,
    // Permissions
    canReadEngineers,
    canWriteEngineers,
    canDeleteEngineers,
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
    resetEngineers,
  }
}

export default useEngineers
