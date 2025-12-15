import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useDebtorsStore } from '@/stores/debtors'
import { usePermissions } from './usePermissions'

/**
 * Composable for debtors management
 * Wraps the debtors store with permission checks
 */
export function useDebtors() {
  const debtorsStore = useDebtorsStore()
  const { canRead, canWrite, canDelete } = usePermissions()

  // Destructure reactive state from store
  const { debtors, loading, error } = storeToRefs(debtorsStore)

  // Permission checks
  const canReadDebtors = canRead('/debtors')
  const canWriteDebtors = canWrite('/debtors')
  const canDeleteDebtors = canDelete('/debtors')

  // Computed statistics
  const totalDebtors = computed(() => debtors.value.length)

  const totalDebt = computed(() => {
    return debtors.value.reduce((sum, debtor) => {
      return sum + (parseFloat(debtor.amount) || 0)
    }, 0)
  })

  const pendingDebtors = computed(() => {
    return debtors.value.filter(d => d.status === 'pending').length
  })

  const paidDebtors = computed(() => {
    return debtors.value.filter(d => d.status === 'paid').length
  })

  const overdueDebtors = computed(() => {
    return debtors.value.filter(d => d.status === 'overdue').length
  })

  // Filtered debtors by status
  const getDebtorsByStatus = (status) => {
    if (!status) return debtors.value
    return debtors.value.filter(d => d.status === status)
  }

  // Search debtors
  const searchDebtors = (query) => {
    if (!query) return debtors.value
    const lowerQuery = query.toLowerCase()
    return debtors.value.filter(d =>
      d.name?.toLowerCase().includes(lowerQuery) ||
      d.email?.toLowerCase().includes(lowerQuery) ||
      d.phone?.includes(query)
    )
  }

  // Actions with permission checks
  const fetchDebtors = async (params = {}) => {
    if (!canReadDebtors.value) {
      console.warn('No permission to read debtors')
      return
    }
    return debtorsStore.fetchDebtors(params)
  }

  const createDebtor = async (data) => {
    if (!canWriteDebtors.value) {
      console.warn('No permission to create debtors')
      return
    }
    return debtorsStore.createDebtor(data)
  }

  const updateDebtor = async (id, data) => {
    if (!canWriteDebtors.value) {
      console.warn('No permission to update debtors')
      return
    }
    return debtorsStore.updateDebtor(id, data)
  }

  const deleteDebtor = async (id) => {
    if (!canDeleteDebtors.value) {
      console.warn('No permission to delete debtors')
      return
    }
    return debtorsStore.deleteDebtor(id)
  }

  const recordPayment = async (id, paymentData) => {
    if (!canWriteDebtors.value) {
      console.warn('No permission to record payments')
      return
    }
    return debtorsStore.recordPayment(id, paymentData)
  }

  return {
    // State
    debtors,
    loading,
    error,

    // Computed statistics
    totalDebtors,
    totalDebt,
    pendingDebtors,
    paidDebtors,
    overdueDebtors,

    // Permissions
    canReadDebtors,
    canWriteDebtors,
    canDeleteDebtors,

    // Methods
    fetchDebtors,
    createDebtor,
    updateDebtor,
    deleteDebtor,
    recordPayment,
    getDebtorsByStatus,
    searchDebtors
  }
}
