import { computed } from 'vue'

/**
 * Composable for project statistics
 * Provides computed statistics and helper functions for project display
 *
 * @param {import('vue').Ref<Array>} projects - Reactive reference to projects array
 * @returns {Object} Statistics and helpers
 */
export function useProjectStats(projects) {
  // Computed statistics
  const totalProjects = computed(() => projects.value?.length || 0)

  const activeProjects = computed(() => projects.value?.filter((p) => p.status === 'active').length || 0)

  const pendingProjects = computed(() => projects.value?.filter((p) => p.status === 'pending').length || 0)

  const completedProjects = computed(() => projects.value?.filter((p) => p.status === 'completed' || p.progress === 100).length || 0)

  const cancelledProjects = computed(() => projects.value?.filter((p) => p.status === 'cancelled').length || 0)

  const averageProgress = computed(() => {
    if (!projects.value || projects.value.length === 0) return 0
    const totalProgress = projects.value.reduce((sum, project) => sum + (project.progress || 0), 0)
    return Math.round(totalProgress / projects.value.length)
  })

  const totalBudget = computed(() => {
    if (!projects.value) return 0
    return projects.value.reduce((sum, project) => sum + (project.initialCost || 0), 0)
  })

  const totalCriticalCost = computed(() => {
    if (!projects.value) return 0
    return projects.value.reduce((sum, project) => sum + (project.criticalCost || 0), 0)
  })

  // Status helpers
  function getStatusColor(status) {
    const colors = {
      pending: 'warning',
      active: 'success',
      completed: 'primary',
      cancelled: 'error',
    }
    return colors[status] || 'grey'
  }

  function getStatusText(status) {
    const texts = {
      pending: 'في الانتظار',
      active: 'نشط',
      completed: 'مكتمل',
      cancelled: 'ملغي',
    }
    return texts[status] || status
  }

  function getStatusIcon(status) {
    const icons = {
      pending: 'mdi-clock-outline',
      active: 'mdi-play-circle',
      completed: 'mdi-check-circle',
      cancelled: 'mdi-cancel',
    }
    return icons[status] || 'mdi-help-circle'
  }

  // Priority helpers
  function getPriorityText(priority) {
    const texts = {
      low: 'منخفضة',
      medium: 'متوسطة',
      high: 'عالية',
      urgent: 'عاجلة',
    }
    return texts[priority] || priority
  }

  function getPriorityColor(priority) {
    const colors = {
      low: 'success',
      medium: 'warning',
      high: 'error',
      urgent: 'purple',
    }
    return colors[priority] || 'grey'
  }

  function getPriorityIcon(priority) {
    const icons = {
      low: 'mdi-arrow-down',
      medium: 'mdi-minus',
      high: 'mdi-arrow-up',
      urgent: 'mdi-alert',
    }
    return icons[priority] || 'mdi-help'
  }

  // Format helpers
  function formatCurrency(amount) {
    if (amount === null || amount === undefined) return '0 د.ع'
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'IQD',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(amount)
  }

  function formatDate(dateString) {
    if (!dateString) return 'غير محدد'
    const date = new Date(dateString)
    return date.toLocaleDateString('en-GB', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    })
  }

  function formatProgress(progress) {
    return `${progress || 0}%`
  }

  // Progress helpers
  function getProgressColor(progress) {
    if (progress >= 80) return 'success'
    if (progress >= 50) return 'primary'
    if (progress >= 25) return 'warning'
    return 'error'
  }

  // Get stats summary object
  const statsSummary = computed(() => ({
    total: totalProjects.value,
    active: activeProjects.value,
    pending: pendingProjects.value,
    completed: completedProjects.value,
    cancelled: cancelledProjects.value,
    averageProgress: averageProgress.value,
    totalBudget: totalBudget.value,
    totalCriticalCost: totalCriticalCost.value,
  }))

  return {
    // Computed statistics
    totalProjects,
    activeProjects,
    pendingProjects,
    completedProjects,
    cancelledProjects,
    averageProgress,
    totalBudget,
    totalCriticalCost,
    statsSummary,
    // Status helpers
    getStatusColor,
    getStatusText,
    getStatusIcon,
    // Priority helpers
    getPriorityText,
    getPriorityColor,
    getPriorityIcon,
    // Format helpers
    formatCurrency,
    formatDate,
    formatProgress,
    getProgressColor,
  }
}

export default useProjectStats
