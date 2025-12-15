import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUiStore = defineStore('ui', () => {
  // Sidebar state
  const sidebarCollapsed = ref(false)
  const sidebarMobileOpen = ref(false)

  // Global loading states
  const globalLoading = ref(false)
  const loadingMessage = ref('')

  // Modal states
  const activeModal = ref(null)
  const modalData = ref(null)

  // Theme
  const darkMode = ref(localStorage.getItem('darkMode') === 'true')

  // Page title
  const pageTitle = ref('')

  // Breadcrumbs
  const breadcrumbs = ref([])

  // Getters
  const isSidebarCollapsed = computed(() => sidebarCollapsed.value)
  const isSidebarMobileOpen = computed(() => sidebarMobileOpen.value)
  const isGlobalLoading = computed(() => globalLoading.value)
  const isDarkMode = computed(() => darkMode.value)
  const hasActiveModal = computed(() => activeModal.value !== null)

  /**
   * Toggle sidebar collapsed state
   */
  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
    localStorage.setItem('sidebarCollapsed', sidebarCollapsed.value)
  }

  /**
   * Set sidebar collapsed state
   * @param {boolean} collapsed - Collapsed state
   */
  function setSidebarCollapsed(collapsed) {
    sidebarCollapsed.value = collapsed
    localStorage.setItem('sidebarCollapsed', collapsed)
  }

  /**
   * Toggle mobile sidebar
   */
  function toggleMobileSidebar() {
    sidebarMobileOpen.value = !sidebarMobileOpen.value
  }

  /**
   * Close mobile sidebar
   */
  function closeMobileSidebar() {
    sidebarMobileOpen.value = false
  }

  /**
   * Open mobile sidebar
   */
  function openMobileSidebar() {
    sidebarMobileOpen.value = true
  }

  /**
   * Show global loading
   * @param {string} [message=''] - Loading message
   */
  function showLoading(message = '') {
    globalLoading.value = true
    loadingMessage.value = message
  }

  /**
   * Hide global loading
   */
  function hideLoading() {
    globalLoading.value = false
    loadingMessage.value = ''
  }

  /**
   * Open a modal
   * @param {string} modalName - Modal identifier
   * @param {*} [data=null] - Data to pass to modal
   */
  function openModal(modalName, data = null) {
    activeModal.value = modalName
    modalData.value = data
  }

  /**
   * Close current modal
   */
  function closeModal() {
    activeModal.value = null
    modalData.value = null
  }

  /**
   * Check if a specific modal is open
   * @param {string} modalName - Modal identifier
   * @returns {boolean}
   */
  function isModalOpen(modalName) {
    return activeModal.value === modalName
  }

  /**
   * Toggle dark mode
   */
  function toggleDarkMode() {
    darkMode.value = !darkMode.value
    localStorage.setItem('darkMode', darkMode.value)
    applyTheme()
  }

  /**
   * Set dark mode
   * @param {boolean} enabled - Dark mode state
   */
  function setDarkMode(enabled) {
    darkMode.value = enabled
    localStorage.setItem('darkMode', enabled)
    applyTheme()
  }

  /**
   * Apply theme to document
   */
  function applyTheme() {
    if (darkMode.value) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }

  /**
   * Set page title
   * @param {string} title - Page title
   */
  function setPageTitle(title) {
    pageTitle.value = title
    document.title = title ? `${title} - نظام إدارة المشاريع` : 'نظام إدارة المشاريع'
  }

  /**
   * Set breadcrumbs
   * @param {Array} items - Breadcrumb items
   */
  function setBreadcrumbs(items) {
    breadcrumbs.value = items
  }

  /**
   * Initialize UI state from localStorage
   */
  function initUi() {
    const savedSidebarState = localStorage.getItem('sidebarCollapsed')
    if (savedSidebarState !== null) {
      sidebarCollapsed.value = savedSidebarState === 'true'
    }

    const savedDarkMode = localStorage.getItem('darkMode')
    if (savedDarkMode !== null) {
      darkMode.value = savedDarkMode === 'true'
    }

    applyTheme()
  }

  /**
   * Reset store state
   */
  function $reset() {
    sidebarCollapsed.value = false
    sidebarMobileOpen.value = false
    globalLoading.value = false
    loadingMessage.value = ''
    activeModal.value = null
    modalData.value = null
    pageTitle.value = ''
    breadcrumbs.value = []
  }

  return {
    // State
    sidebarCollapsed,
    sidebarMobileOpen,
    globalLoading,
    loadingMessage,
    activeModal,
    modalData,
    darkMode,
    pageTitle,
    breadcrumbs,
    // Getters
    isSidebarCollapsed,
    isSidebarMobileOpen,
    isGlobalLoading,
    isDarkMode,
    hasActiveModal,
    // Actions
    toggleSidebar,
    setSidebarCollapsed,
    toggleMobileSidebar,
    closeMobileSidebar,
    openMobileSidebar,
    showLoading,
    hideLoading,
    openModal,
    closeModal,
    isModalOpen,
    toggleDarkMode,
    setDarkMode,
    applyTheme,
    setPageTitle,
    setBreadcrumbs,
    initUi,
    $reset,
  }
})
