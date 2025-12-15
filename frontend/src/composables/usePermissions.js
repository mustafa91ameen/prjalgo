import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useAuthStore } from '@/stores/auth'

/**
 * Composable for checking user permissions
 * Provides reactive permission checking based on current user's pages/roles
 *
 * @returns {Object} Permission checking utilities
 */
export function usePermissions() {
  const authStore = useAuthStore()
  const { pages, isAuthenticated } = storeToRefs(authStore)

  /**
   * Check if user has permission for a specific page and action
   * @param {string} route - Page route (e.g., '/projects')
   * @param {string} permission - Permission type ('read', 'write', 'delete')
   * @returns {boolean}
   */
  function hasPermission(route, permission) {
    if (!isAuthenticated.value) return false

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
   * Check if user can write to a page (create/update)
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
   * Check if user has access to a page (can read)
   * @param {string} route - Page route
   * @returns {boolean}
   */
  function hasAccess(route) {
    return canRead(route)
  }

  /**
   * Get all permissions for a specific page
   * @param {string} route - Page route
   * @returns {Object} Permissions object with read, write, delete booleans
   */
  function getPagePermissions(route) {
    return {
      read: canRead(route),
      write: canWrite(route),
      delete: canDelete(route),
    }
  }

  /**
   * Get all accessible page routes
   * @returns {string[]} Array of accessible routes
   */
  const accessibleRoutes = computed(() => {
    return pages.value.filter((p) => p.permissions?.includes('read')).map((p) => p.route)
  })

  /**
   * Check if user has any of the specified permissions
   * @param {string} route - Page route
   * @param {string[]} permissions - Array of permission types
   * @returns {boolean}
   */
  function hasAnyPermission(route, permissions) {
    return permissions.some((permission) => hasPermission(route, permission))
  }

  /**
   * Check if user has all of the specified permissions
   * @param {string} route - Page route
   * @param {string[]} permissions - Array of permission types
   * @returns {boolean}
   */
  function hasAllPermissions(route, permissions) {
    return permissions.every((permission) => hasPermission(route, permission))
  }

  return {
    // Methods
    hasPermission,
    canRead,
    canWrite,
    canDelete,
    hasAccess,
    getPagePermissions,
    hasAnyPermission,
    hasAllPermissions,
    // Computed
    accessibleRoutes,
    // State refs (from store)
    pages,
    isAuthenticated,
  }
}

export default usePermissions
