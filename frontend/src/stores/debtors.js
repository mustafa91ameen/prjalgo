import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import { debtorsApi } from '@/api/debtors'

export const useDebtorsStore = defineStore('debtors', () => {
  // State
  const debtors = ref([])
  const currentDebtor = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0,
    totalPages: 0,
  })

  // Getters
  const debtorCount = computed(() => debtors.value.length)
  const totalDebtors = computed(() => pagination.value.total)
  const hasDebtors = computed(() => debtors.value.length > 0)
  const isLoading = computed(() => loading.value)
  const totalDebt = computed(() => {
    return debtors.value.reduce((sum, debtor) => sum + (debtor.amount || 0), 0)
  })

  /**
   * Fetch all debtors with pagination
   * @param {Object} params - Query parameters
   */
  async function fetchDebtors(params = {}) {
    loading.value = true
    error.value = null

    try {
      const response = await debtorsApi.getAll({
        page: params.page || pagination.value.page,
        limit: params.limit || pagination.value.limit,
        ...params,
      })

      debtors.value = response.data ?? []
      pagination.value = {
        page: response.page ?? 1,
        limit: response.limit ?? 10,
        total: response.total ?? 0,
        totalPages: response.totalPages ?? 0,
      }
    } catch (err) {
      error.value = err.message
      debtors.value = []
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch a single debtor by ID
   * @param {number} id - Debtor ID
   */
  async function fetchDebtorById(id) {
    loading.value = true
    error.value = null

    try {
      const response = await debtorsApi.getById(id)
      currentDebtor.value = response.data ?? response
      return currentDebtor.value
    } catch (err) {
      error.value = err.message
      currentDebtor.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Create a new debtor
   * @param {Object} data - Debtor data
   * @returns {Promise<Object|null>} - Created debtor or null
   */
  async function createDebtor(data) {
    loading.value = true
    error.value = null

    try {
      const response = await debtorsApi.create(data)
      const newDebtor = response.data ?? response
      debtors.value.unshift(newDebtor)
      pagination.value.total += 1
      toast.success('تم إضافة المدين بنجاح')
      return newDebtor
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Update an existing debtor
   * @param {number} id - Debtor ID
   * @param {Object} data - Updated debtor data
   * @returns {Promise<Object|null>} - Updated debtor or null
   */
  async function updateDebtor(id, data) {
    loading.value = true
    error.value = null

    try {
      const response = await debtorsApi.update(id, data)
      const updatedDebtor = response.data ?? response

      const index = debtors.value.findIndex((d) => d.id === id)
      if (index !== -1) {
        debtors.value[index] = updatedDebtor
      }

      if (currentDebtor.value?.id === id) {
        currentDebtor.value = updatedDebtor
      }

      toast.success('تم تحديث المدين بنجاح')
      return updatedDebtor
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Delete a debtor
   * @param {number} id - Debtor ID
   * @returns {Promise<boolean>} - Success status
   */
  async function deleteDebtor(id) {
    loading.value = true
    error.value = null

    try {
      await debtorsApi.delete(id)

      debtors.value = debtors.value.filter((d) => d.id !== id)
      pagination.value.total -= 1

      if (currentDebtor.value?.id === id) {
        currentDebtor.value = null
      }

      toast.success('تم حذف المدين بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
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
   * Clear current debtor
   */
  function clearCurrentDebtor() {
    currentDebtor.value = null
  }

  /**
   * Reset store state
   */
  function $reset() {
    debtors.value = []
    currentDebtor.value = null
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
    debtors,
    currentDebtor,
    loading,
    error,
    pagination,
    // Getters
    debtorCount,
    totalDebtors,
    hasDebtors,
    isLoading,
    totalDebt,
    // Actions
    fetchDebtors,
    fetchDebtorById,
    createDebtor,
    updateDebtor,
    deleteDebtor,
    setPage,
    setLimit,
    clearCurrentDebtor,
    $reset,
  }
})
