import { ref, readonly } from 'vue'

// Global toast state (shared across all components)
const toast = ref({
  show: false,
  message: '',
  color: 'success' // success, error, warning, info
})

let timeoutId = null

export function useToast() {
  const showToast = (message, color = 'success') => {
    // Clear any existing timeout
    if (timeoutId) {
      clearTimeout(timeoutId)
    }

    toast.value = { show: true, message, color }

    // Auto-hide after 3 seconds
    timeoutId = setTimeout(() => {
      toast.value.show = false
    }, 3000)
  }

  const hideToast = () => {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }
    toast.value.show = false
  }

  const success = (message) => showToast(message, 'success')
  const error = (message) => showToast(message, 'error')
  const warning = (message) => showToast(message, 'warning')
  const info = (message) => showToast(message, 'info')

  return {
    toast: readonly(toast),
    showToast,
    hideToast,
    success,
    error,
    warning,
    info
  }
}

// Get toast icon based on type
export function getToastIcon(color) {
  const icons = {
    success: 'mdi mdi-check-circle',
    error: 'mdi mdi-alert-circle',
    warning: 'mdi mdi-alert',
    info: 'mdi mdi-information'
  }
  return icons[color] || icons.info
}
