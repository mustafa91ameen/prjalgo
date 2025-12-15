import { storeToRefs } from 'pinia'
import { useExpensesStore } from '@/stores/expenses'
import { usePermissions } from './usePermissions'

/**
 * Composable for expense-related logic
 * Wraps the expenses store with permission checks and additional utilities
 *
 * @returns {Object} Expense utilities
 */
export function useExpenses() {
  const expensesStore = useExpensesStore()
  const { canRead, canWrite, canDelete } = usePermissions()

  // Get reactive state from store using storeToRefs
  const {
    expenses,
    currentExpense,
    loading,
    error,
    pagination,
    filters,
    expenseCount,
    totalExpenses,
    hasExpenses,
    isLoading,
    totalAmount,
  } = storeToRefs(expensesStore)

  // Permission checks for expenses page
  const canReadExpenses = canRead('/expenses')
  const canWriteExpenses = canWrite('/expenses')
  const canDeleteExpenses = canDelete('/expenses')

  /**
   * Fetch all expenses (with permission check)
   * @param {Object} params - Query parameters
   */
  async function fetchExpenses(params = {}) {
    if (!canReadExpenses) {
      return
    }
    return expensesStore.fetchExpenses(params)
  }

  /**
   * Fetch a single expense by ID
   * @param {number} id - Expense ID
   */
  async function fetchExpenseById(id) {
    if (!canReadExpenses) {
      return null
    }
    return expensesStore.fetchExpenseById(id)
  }

  /**
   * Create a new expense (with permission check)
   * @param {Object} data - Expense data
   */
  async function createExpense(data) {
    if (!canWriteExpenses) {
      return null
    }
    return expensesStore.createExpense(data)
  }

  /**
   * Update an expense (with permission check)
   * @param {number} id - Expense ID
   * @param {Object} data - Updated data
   */
  async function updateExpense(id, data) {
    if (!canWriteExpenses) {
      return null
    }
    return expensesStore.updateExpense(id, data)
  }

  /**
   * Delete an expense (with permission check)
   * @param {number} id - Expense ID
   */
  async function deleteExpense(id) {
    if (!canDeleteExpenses) {
      return false
    }
    return expensesStore.deleteExpense(id)
  }

  /**
   * Set filters
   * @param {Object} newFilters - Filter values
   */
  function setFilters(newFilters) {
    expensesStore.setFilters(newFilters)
  }

  /**
   * Clear filters
   */
  function clearFilters() {
    expensesStore.clearFilters()
  }

  /**
   * Set current page for pagination
   * @param {number} page - Page number
   */
  function setPage(page) {
    expensesStore.setPage(page)
  }

  /**
   * Set items per page
   * @param {number} limit - Items per page
   */
  function setLimit(limit) {
    expensesStore.setLimit(limit)
  }

  /**
   * Clear current expense
   */
  function clearCurrentExpense() {
    expensesStore.clearCurrentExpense()
  }

  /**
   * Reset store state
   */
  function resetExpenses() {
    expensesStore.$reset()
  }

  return {
    // State (reactive refs)
    expenses,
    currentExpense,
    loading,
    error,
    pagination,
    filters,
    // Getters (reactive)
    expenseCount,
    totalExpenses,
    hasExpenses,
    isLoading,
    totalAmount,
    // Permissions
    canReadExpenses,
    canWriteExpenses,
    canDeleteExpenses,
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
    resetExpenses,
  }
}

export default useExpenses
