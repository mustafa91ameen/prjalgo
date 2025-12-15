import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { toast } from 'vue3-toastify'
import { expensesApi } from '@/api/expenses'

export const useExpensesStore = defineStore('expenses', () => {
  // State
  const expenses = ref([])
  const currentExpense = ref(null)
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    page: 1,
    limit: 10,
    total: 0,
    totalPages: 0,
  })
  const filters = ref({
    projectId: null,
    startDate: null,
    endDate: null,
    search: '',
  })

  // Getters
  const expenseCount = computed(() => expenses.value.length)
  const totalExpenses = computed(() => pagination.value.total)
  const hasExpenses = computed(() => expenses.value.length > 0)
  const isLoading = computed(() => loading.value)
  const totalAmount = computed(() => {
    return expenses.value.reduce((sum, expense) => sum + (expense.amount || 0), 0)
  })

  /**
   * Fetch all expenses with pagination and filters
   * @param {Object} params - Query parameters
   */
  async function fetchExpenses(params = {}) {
    loading.value = true
    error.value = null

    try {
      const response = await expensesApi.getAll({
        page: params.page || pagination.value.page,
        limit: params.limit || pagination.value.limit,
        projectId: params.projectId ?? filters.value.projectId,
        startDate: params.startDate ?? filters.value.startDate,
        endDate: params.endDate ?? filters.value.endDate,
        search: params.search ?? filters.value.search,
        ...params,
      })

      expenses.value = response.data ?? []
      pagination.value = {
        page: response.page ?? 1,
        limit: response.limit ?? 10,
        total: response.total ?? 0,
        totalPages: response.totalPages ?? 0,
      }
    } catch (err) {
      error.value = err.message
      expenses.value = []
    } finally {
      loading.value = false
    }
  }

  /**
   * Fetch a single expense by ID
   * @param {number} id - Expense ID
   */
  async function fetchExpenseById(id) {
    loading.value = true
    error.value = null

    try {
      const response = await expensesApi.getById(id)
      currentExpense.value = response.data ?? response
      return currentExpense.value
    } catch (err) {
      error.value = err.message
      currentExpense.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Create a new expense
   * @param {Object} data - Expense data
   * @returns {Promise<Object|null>} - Created expense or null
   */
  async function createExpense(data) {
    loading.value = true
    error.value = null

    try {
      const response = await expensesApi.create(data)
      const newExpense = response.data ?? response
      expenses.value.unshift(newExpense)
      pagination.value.total += 1
      toast.success('تم إضافة المصروف بنجاح')
      return newExpense
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Update an existing expense
   * @param {number} id - Expense ID
   * @param {Object} data - Updated expense data
   * @returns {Promise<Object|null>} - Updated expense or null
   */
  async function updateExpense(id, data) {
    loading.value = true
    error.value = null

    try {
      const response = await expensesApi.update(id, data)
      const updatedExpense = response.data ?? response

      const index = expenses.value.findIndex((e) => e.id === id)
      if (index !== -1) {
        expenses.value[index] = updatedExpense
      }

      if (currentExpense.value?.id === id) {
        currentExpense.value = updatedExpense
      }

      toast.success('تم تحديث المصروف بنجاح')
      return updatedExpense
    } catch (err) {
      error.value = err.message
      return null
    } finally {
      loading.value = false
    }
  }

  /**
   * Delete an expense
   * @param {number} id - Expense ID
   * @returns {Promise<boolean>} - Success status
   */
  async function deleteExpense(id) {
    loading.value = true
    error.value = null

    try {
      await expensesApi.delete(id)

      expenses.value = expenses.value.filter((e) => e.id !== id)
      pagination.value.total -= 1

      if (currentExpense.value?.id === id) {
        currentExpense.value = null
      }

      toast.success('تم حذف المصروف بنجاح')
      return true
    } catch (err) {
      error.value = err.message
      return false
    } finally {
      loading.value = false
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
      projectId: null,
      startDate: null,
      endDate: null,
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
   * Clear current expense
   */
  function clearCurrentExpense() {
    currentExpense.value = null
  }

  /**
   * Reset store state
   */
  function $reset() {
    expenses.value = []
    currentExpense.value = null
    loading.value = false
    error.value = null
    pagination.value = {
      page: 1,
      limit: 10,
      total: 0,
      totalPages: 0,
    }
    filters.value = {
      projectId: null,
      startDate: null,
      endDate: null,
      search: '',
    }
  }

  return {
    // State
    expenses,
    currentExpense,
    loading,
    error,
    pagination,
    filters,
    // Getters
    expenseCount,
    totalExpenses,
    hasExpenses,
    isLoading,
    totalAmount,
    // Actions
    fetchExpenses,
    fetchExpenseById,
    createExpense,
    updateExpense,
    deleteExpense,
    setFilters,
    clearFilters,
    setPage,
    setLimit,
    clearCurrentExpense,
    $reset,
  }
})
