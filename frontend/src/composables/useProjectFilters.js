import { ref, computed } from 'vue'

/**
 * Composable for project filtering and search logic
 * Extracts filter state and computed filtering logic from project-management.vue
 *
 * @param {import('vue').Ref<Array>} projects - Reactive reference to projects array
 * @returns {Object} Filter utilities
 */
export function useProjectFilters(projects) {
  // Filter state
  const searchQuery = ref('')
  const selectedCategory = ref('')
  const selectedStatus = ref('')

  // Filter options
  const filterCategories = [
    { title: 'تطوير', value: 'تطوير' },
    { title: 'تحديث', value: 'تحديث' },
    { title: 'بناء', value: 'بناء' },
    { title: 'صيانة', value: 'صيانة' },
    { title: 'استشارات', value: 'استشارات' },
  ]

  const filterStatuses = [
    { title: 'في الانتظار', value: 'pending' },
    { title: 'نشط', value: 'active' },
    { title: 'مكتمل', value: 'completed' },
    { title: 'ملغي', value: 'cancelled' },
  ]

  // Computed filtered projects
  const filteredProjects = computed(() => {
    let filtered = projects.value || []

    // Filter by search query
    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      filtered = filtered.filter(
        (project) =>
          project.name?.toLowerCase().includes(query) ||
          project.location?.toLowerCase().includes(query) ||
          project.user?.toLowerCase().includes(query) ||
          project.description?.toLowerCase().includes(query) ||
          project.category?.toLowerCase().includes(query)
      )
    }

    // Filter by category
    if (selectedCategory.value) {
      filtered = filtered.filter((project) => project.category === selectedCategory.value)
    }

    // Filter by status
    if (selectedStatus.value) {
      filtered = filtered.filter((project) => project.status === selectedStatus.value)
    }

    return filtered
  })

  // Reset all filters
  function resetFilters() {
    searchQuery.value = ''
    selectedCategory.value = ''
    selectedStatus.value = ''
  }

  // Check if any filter is active
  const hasActiveFilters = computed(() => {
    return Boolean(searchQuery.value || selectedCategory.value || selectedStatus.value)
  })

  return {
    // State
    searchQuery,
    selectedCategory,
    selectedStatus,
    // Options
    filterCategories,
    filterStatuses,
    // Computed
    filteredProjects,
    hasActiveFilters,
    // Actions
    resetFilters,
  }
}

export default useProjectFilters
